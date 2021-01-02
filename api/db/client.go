package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://admin:root@mongodb:27017"))
	if err != nil {
		fmt.Printf("client:%v\n", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Printf("conn:%v\n", err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Printf("ping:%v\n", err)
	}
}
