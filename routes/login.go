package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/GonzaPiccinini/twitter-go/consts"
	"github.com/GonzaPiccinini/twitter-go/db"
	"github.com/GonzaPiccinini/twitter-go/jwt"
	"github.com/GonzaPiccinini/twitter-go/models"
	"github.com/aws/aws-lambda-go/events"
)

func Login(ctx context.Context) models.APIResponse {
	var u models.User
	var r models.APIResponse
	r.Status = 400

	fmt.Println("-> Login handler")

	body := ctx.Value(models.Key(consts.BODY)).(string)
	err := json.Unmarshal([]byte(body), &u)
	if err != nil {
		fmt.Println("Error unmarshaling body")
		r.Message = "invalid email or password"
		return r
	}
	if len(u.Email) == 0 {
		fmt.Println("Invalid email")
		r.Message = "invalid email or password"
		return r
	}

	user, ok := db.Login(u.Email, u.Password)
	if !ok {
		r.Message = "invalid email or password"
		return r
	}

	token, err := jwt.GenerateJWTToken(ctx, user)
	if err != nil {
		fmt.Println("Error generating token")
		r.Message = "invalid email or password"
		return r
	}

	loginResponse := models.LoginResponse{
		Token: token,
	}

	tokenMarshaled, err := json.Marshal(loginResponse)
	if err != nil {
		fmt.Println("Error marshaling token")
		r.Message = "invalid email or password"
		return r
	}

	cookie := &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 24),
	}

	cookieStr := cookie.String()

	cookieResponse := &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(tokenMarshaled),
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
			"Set-Cookie":                  cookieStr,
		},
	}

	r.Status = 200
	r.Message = string(tokenMarshaled)
	r.Response = cookieResponse
	return r
}
