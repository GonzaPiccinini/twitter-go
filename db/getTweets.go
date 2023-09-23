package db

import (
	"context"

	"github.com/GonzaPiccinini/twitter-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetTweets(ID string, page int64) ([]*models.GetTweets, bool) {
	ctx := context.TODO()

	database := MONGO_CNN.Database(Database)
	collection := database.Collection("tweets")

	var results []*models.GetTweets

	condition := bson.M{
		"userid": ID,
	}
	optionss := options.Find()
	optionss.SetLimit(20)
	optionss.SetSort(bson.D{{Key: "date", Value: -1}})
	optionss.SetSkip((page - 1) * 20)

	cursor, err := collection.Find(ctx, condition, optionss)
	if err != nil {
		return results, false
	}

	for cursor.Next(ctx) {
		var doc models.GetTweets
		err := cursor.Decode(&doc)
		if err != nil {
			return results, false
		}
		results = append(results, &doc)
	}

	return results, true
}
