package routes

import (
	"github.com/GonzaPiccinini/twitter-go/db"
	"github.com/GonzaPiccinini/twitter-go/models"
	"github.com/aws/aws-lambda-go/events"
)

func DeleteTweet(request events.APIGatewayProxyRequest, claims models.Claim) models.APIResponse {
	var r models.APIResponse
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "id param is required"
		return r
	}

	err := db.DeleteTweet(ID, claims.ID.Hex())
	if err != nil {
		r.Message = "Error deleting tweet"
		return r
	}

	r.Message = "Successful tweet deletion"
	r.Status = 200
	return r
}
