package routes

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/GonzaPiccinini/twitter-go/models"
)

func Register(ctx context.Context) models.APIResponse {
	var u models.User
	var r models.APIResponse
	r.Status = 400

	fmt.Println("-> Register handler")

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &u)
	if err != nil {
		r.Message = err.Error()
		fmt.Println("Error unmarshaling body" + r.Message)
		return r
	}

	if len(u.Email) == 0 {
		r.Message = "invalid email or password"
		fmt.Println(r.Message)
		return r
	}
	if len(u.Email) < 6 {
		r.Message = "the length of the password must be greater than 6"
		fmt.Println(r.Message)
		return r
	}

	_, found, _ := db.UserExists(u.Email)
	if found {
		r.Message = "user already exists"
		fmt.Println(r.Message)
		return r
	}

	_, status, err := db.Register(u)
	if err != nil {
		r.Message = "error when trying to register the user" + err.Error()
		fmt.Println(r.Message)
		return r
	}
	if !status {
		r.Message = "error when trying to register the user"
		fmt.Println(r.Message)
		return r
	}

	r.Status = 201
	r.Message = "successful register"
	fmt.Println("successful register")
	return r
}
