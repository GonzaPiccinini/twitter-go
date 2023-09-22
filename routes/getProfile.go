package routes

import (
	"encoding/json"
	"fmt"

	"github.com/GonzaPiccinini/twitter-go/db"
	"github.com/GonzaPiccinini/twitter-go/models"
	"github.com/aws/aws-lambda-go/events"
)

func GetProfile(request events.APIGatewayProxyRequest) models.APIResponse {
	var r models.APIResponse
	r.Status = 400

	fmt.Println("-> GetProfile handler")

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "ID parameter is required"
		return r
	}

	profile, err := db.GetProfile(ID)
	if err != nil {
		r.Message = "Error finding profile"
		return r
	}

	profileMarshaled, err := json.Marshal(profile)
	if err != nil {
		r.Status = 500
		r.Message = "Error marshaling profile"
		return r
	}

	r.Status = 200
	r.Message = string(profileMarshaled)
	return r
}
