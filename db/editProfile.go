package db

import (
	"context"

	"github.com/GonzaPiccinini/twitter-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func EditProfile(u models.User, ID string) (bool, error) {
	ctx := context.TODO()

	database := MONGO_CNN.Database(Database)
	collection := database.Collection("users")

	document := make(map[string]interface{})
	if len(u.Firstname) > 0 {
		document["firstname"] = u.Firstname
	}
	if len(u.Lastname) > 0 {
		document["lastname"] = u.Lastname
	}
	document["birthdate"] = u.Birthdate
	if len(u.Avatar) > 0 {
		document["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		document["banner"] = u.Banner
	}
	if len(u.Biography) > 0 {
		document["biography"] = u.Biography
	}
	if len(u.Web) > 0 {
		document["web"] = u.Web
	}
	if len(u.Ubication) > 0 {
		document["ubication"] = u.Ubication
	}

	update := bson.M{
		"$set": document,
	}
	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{
		"_id": bson.M{"$eq": objID},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}

	return true, nil
}
