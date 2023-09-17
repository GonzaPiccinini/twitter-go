package models

import (
	"github.com/aws/aws-lambda-go/events"
)

type APIResponse struct {
	Status   int
	Message  string
	Response *events.APIGatewayProxyResponse
}
