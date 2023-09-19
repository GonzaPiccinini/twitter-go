package db

import (
	"context"

	"github.com/GonzaPiccinini/twitter-go/models"
	"go.mongodb.org/mongo-driver/bson"
)

func UserExists(email string) (models.User, bool, string) {
	var u models.User

	ctx := context.TODO()
	database := MONGO_CNN.Database(Database)
	collection := database.Collection("users")

	condition := bson.M{"email": email}

	err := collection.FindOne(ctx, condition).Decode(&u)
	ID := u.ID.Hex()
	if err != nil {
		return u, false, ID
	}
	return u, true, ID
}
