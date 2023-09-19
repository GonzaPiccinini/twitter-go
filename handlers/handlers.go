package handlers

import (
	"context"
	"fmt"

	"github.com/GonzaPiccinini/twitter-go/jwt"
	"github.com/GonzaPiccinini/twitter-go/models"
	"github.com/GonzaPiccinini/twitter-go/routes"
	"github.com/aws/aws-lambda-go/events"
)

func Handlers(ctx context.Context, request events.APIGatewayProxyRequest) models.APIResponse {
	fmt.Println("-> Processing " + ctx.Value(models.Key("path")).(string) + " (" + ctx.Value(models.Key("method")).(string) + ")")

	var res models.APIResponse
	res.Status = 400

	ok, statusCode, message, _ := validateAuthorization(ctx, request)
	if !ok {
		res.Status = statusCode
		res.Message = message
		return res
	}

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {
		case "register":
			return routes.Register(ctx)
		}
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {

		}
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {

		}
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {

		}
	}

	res.Message = "invalid method"
	return res
}

func validateAuthorization(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	path := ctx.Value(models.Key("path")).(string)

	if path == "register" || path == "login" || path == "getAvatar" || path == "getBanner" {
		return true, 200, "", models.Claim{}
	}

	fmt.Println("-> Processing token")

	token := request.Headers["Authorization"]
	if len(token) == 0 {
		return false, 400, "invalid token", models.Claim{}
	}
	claims, ok, message, err := jwt.ProcessToken(token, ctx.Value(models.Key("jwtsign")).(string))
	if !ok {
		if err != nil {
			fmt.Println("Error processing JWT token" + err.Error())
			return false, 401, err.Error(), models.Claim{}
		} else {
			fmt.Println("Error processing JWT token" + message)
			return false, 401, message, models.Claim{}
		}
	}

	fmt.Println("-> Successful token")
	return true, 200, message, *claims
}
