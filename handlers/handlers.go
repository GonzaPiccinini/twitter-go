package handlers

import (
	"context"
	"fmt"

	"github.com/GonzaPiccinini/twitter-go/models"
	"github.com/aws/aws-lambda-go/events"
)

func Handlers(ctx context.Context, request events.APIGatewayProxyRequest) models.APIResponse {
	fmt.Println("-> Processing " + ctx.Value(models.Key("path")).(string) + " (" + ctx.Value(models.Key("method")).(string) + ")")

	var res models.APIResponse
	res.Status = 400

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {

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
