package db

import (
	"context"

	"github.com/GonzaPiccinini/twitter-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTweet(tweet models.CreateTweet) (string, bool, error) {
	ctx := context.TODO()

	database := MONGO_CNN.Database(Database)
	collection := database.Collection("tweets")

	document := bson.M{
		"userid":  tweet.UserID,
		"message": tweet.Message,
		"date":    tweet.Date,
	}

	result, err := collection.InsertOne(ctx, document)
	if err != nil {
		return err.Error(), false, err
	}
	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}
