package db

import (
	"context"

	"github.com/GonzaPiccinini/twitter-go/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Register(u models.User) (string, bool, error) {
	ctx := context.TODO()

	database := MONGO_CNN.Database(Database)
	collection := database.Collection("users")

	u.Password, _ = EncryptPassword(u.Password)

	result, err := collection.InsertOne(ctx, u)
	if err != nil {
		return err.Error(), false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
