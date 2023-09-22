package routes

import (
	"context"
	"encoding/json"

	"github.com/GonzaPiccinini/twitter-go/consts"
	"github.com/GonzaPiccinini/twitter-go/db"
	"github.com/GonzaPiccinini/twitter-go/models"
)

func EditProfile(ctx context.Context, claims models.Claim) models.APIResponse {
	var u models.User
	var r models.APIResponse
	r.Status = 400

	body := ctx.Value(models.Key(consts.BODY)).(string)
	err := json.Unmarshal([]byte(body), &u)
	if err != nil {
		r.Message = "invalid data"
		return r
	}

	ok, err := db.EditProfile(u, claims.ID.Hex())
	if err != nil {
		r.Message = "Error editing profile"
		return r
	}
	if !ok {
		r.Message = "Error editing profile"
		return r
	}

	r.Status = 200
	r.Message = "Successful profile edition"
	return r
}
