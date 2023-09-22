package db

import (
	"context"

	"github.com/GonzaPiccinini/twitter-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetProfile(ID string) (models.User, error) {
	ctx := context.TODO()
	database := MONGO_CNN.Database(Database)
	collection := database.Collection("users")

	var profile models.User

	objID, _ := primitive.ObjectIDFromHex(ID)
	condition := bson.M{"_id": objID}
	err := collection.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""
	if err != nil {
		return profile, err
	}

	return profile, nil
}
