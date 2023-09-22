package routes

import (
	"encoding/json"
	"strconv"

	"github.com/GonzaPiccinini/twitter-go/db"
	"github.com/GonzaPiccinini/twitter-go/models"
	"github.com/aws/aws-lambda-go/events"
)

func GetTweets(request events.APIGatewayProxyRequest) models.APIResponse {
	var r models.APIResponse
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	page := request.QueryStringParameters["pagina"]

	if len(ID) < 1 {
		r.Message = "id param is required"
		return r
	}
	if len(page) < 1 {
		page = "1"
	}

	pagInt, err := strconv.Atoi(page)
	if err != nil {
		r.Message = "page param must be greater than 0"
		return r
	}

	tweets, ok := db.GetTweets(ID, int64(pagInt))
	if !ok {
		r.Message = "error reading tweets"
		return r
	}

	tweetsMarhsaled, err := json.Marshal(tweets)
	if err != nil {
		r.Status = 500
		r.Message = "errro marshaling tweets"
		return r
	}

	r.Status = 200
	r.Message = string(tweetsMarhsaled)
	return r
}
