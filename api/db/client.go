package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	client *mongo.Client
)

// Connect attempts to connect to the mongodb instance.
func Connect(user, pw string) (err error) {
	uri := fmt.Sprintf("mongodb://%s:%s@mongodb:27017", user, pw)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return
	}
	err = ping(ctx)
	return
}

// Disconnect should be called when the client is no longer needed.
func Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
}

func ping(ctx context.Context) (err error) {
	for i := 0; i < 10; i++ {
		err = client.Ping(ctx, readpref.Primary())
		if err == nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
	return
}
