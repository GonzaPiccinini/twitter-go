package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteTweet(ID string, UserID string) error {
	ctx := context.TODO()

	database := MONGO_CNN.Database(Database)
	collection := database.Collection("tweets")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id":    objID,
		"userid": UserID,
	}

	_, err := collection.DeleteOne(ctx, condition)

	return err
}
