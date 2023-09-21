package routes

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/GonzaPiccinini/twitter-go/consts"
	"github.com/GonzaPiccinini/twitter-go/db"
	"github.com/GonzaPiccinini/twitter-go/models"
)

func Register(ctx context.Context) models.APIResponse {
	var u models.User
	var r models.APIResponse
	r.Status = 400

	fmt.Println("-> Register handler")

	body := ctx.Value(models.Key(consts.BODY)).(string)
	err := json.Unmarshal([]byte(body), &u)
	if err != nil {
		r.Message = err.Error()
		return r
	}

	if len(u.Email) == 0 {
		r.Message = "invalid email or password"
		return r
	}
	if len(u.Email) < 6 {
		r.Message = "the length of the password must be greater than 6"
		return r
	}

	_, found, _ := db.UserExists(u.Email)
	if found {
		r.Message = "user already exists"
		return r
	}

	_, status, err := db.Register(u)
	if err != nil {
		r.Message = "error when trying to register the user" + err.Error()
		return r
	}
	if !status {
		r.Message = "error when trying to register the user"
		return r
	}

	r.Status = 201
	r.Message = "successful register"
	fmt.Println("successful register")
	return r
}
