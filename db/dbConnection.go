package db

import (
	"context"
	"fmt"

	"github.com/GonzaPiccinini/twitter-go/consts"
	"github.com/GonzaPiccinini/twitter-go/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MONGO_CNN *mongo.Client
var Database string

func ConnectDB(ctx context.Context) error {
	user := ctx.Value(models.Key(consts.USER)).(string)
	password := ctx.Value(models.Key(consts.PASSWORD)).(string)
	host := ctx.Value(models.Key(consts.HOST)).(string)

	fmt.Println("-> Connecting to Mongo Database")

	connStr := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", user, password, host)

	clientOptions := options.Client().ApplyURI(connStr)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("-> Successful database connection")

	MONGO_CNN = client
	Database = ctx.Value(models.Key(consts.DB_COLLECTION)).(string)

	return nil
}

func PingConnection() bool {
	err := MONGO_CNN.Ping(context.TODO(), nil)
	return err == nil
}
