package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/taivama/golang-training/constants"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDB(ctx context.Context) (*mongo.Client, error) {
	conn := options.Client().ApplyURI(constants.DBConnection)
	client, err := mongo.Connect(ctx, conn)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("ping to mongodb failed")
		return nil, err
	}
	fmt.Println("connected to MongoDB")
	return client, nil
}
