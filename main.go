package main

import (
	"context"
	"os"
	"strings"

	"github.com/GonzaPiccinini/twitter-go/awsgo"
	"github.com/GonzaPiccinini/twitter-go/consts"
	"github.com/GonzaPiccinini/twitter-go/db"
	"github.com/GonzaPiccinini/twitter-go/handlers"
	"github.com/GonzaPiccinini/twitter-go/models"
	"github.com/GonzaPiccinini/twitter-go/secretsmanager"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(runLambda)
}

func runLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse

	awsgo.InitializeAws()

	if !validateParams() {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error in environment variables. Must include ['SecretName', 'BucketName', 'UrlPrefix']",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	secret, err := secretsmanager.GetSecret(os.Getenv(consts.SECRET_NAME))
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error reading secrets: " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	path := strings.Replace(request.PathParameters["twittergo"], os.Getenv(consts.URL_PREFIX), "", -1)

	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key(consts.PATH), path)

	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key(consts.METHOD), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key(consts.BODY), request.Body)

	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key(consts.USER), secret.Username)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key(consts.PASSWORD), secret.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key(consts.HOST), secret.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key(consts.DB_COLLECTION), secret.DbCollection)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key(consts.JWTSIGN), secret.JWTSign)

	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key(consts.BUCKET_NAME), os.Getenv(consts.BUCKET_NAME))

	// Connecting to database
	err = db.ConnectDB(awsgo.Ctx)
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error connecting to database: " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	resApi := handlers.Handlers(awsgo.Ctx, request)
	if resApi.Response == nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: resApi.Status,
			Body:       resApi.Message,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	} else {
		return resApi.Response, nil
	}
}

func validateParams() bool {
	_, ok := os.LookupEnv(consts.SECRET_NAME)
	if !ok {
		return ok
	}

	_, ok = os.LookupEnv(consts.BUCKET_NAME)
	if !ok {
		return ok
	}

	_, ok = os.LookupEnv(consts.URL_PREFIX)
	if !ok {
		return ok
	}

	return ok
}
