package routes

import (
	"context"
	"encoding/json"
	"time"

	"github.com/GonzaPiccinini/twitter-go/consts"
	"github.com/GonzaPiccinini/twitter-go/db"
	"github.com/GonzaPiccinini/twitter-go/models"
)

func CreateTweet(ctx context.Context, claims models.Claim) models.APIResponse {
	var tweet models.Tweet
	var r models.APIResponse
	r.Status = 400

	ID := claims.ID.Hex()
	body := ctx.Value(models.Key(consts.BODY)).(string)
	err := json.Unmarshal([]byte(body), &tweet)
	if err != nil {
		r.Message = "Error unmarshaling data"
		return r
	}

	register := models.CreateTweet{
		UserID:  ID,
		Message: tweet.Message,
		Date:    time.Now(),
	}

	_, ok, err := db.CreateTweet(register)
	if err != nil {
		r.Message = "Error creating tweet"
		return r
	}
	if !ok {
		r.Message = "Error creating tweet"
		return r
	}

	r.Status = 201
	r.Message = "Successful tweet creation"
	return r
}
