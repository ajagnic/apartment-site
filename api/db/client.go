package db

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	client *mongo.Client
	coll   *mongo.Collection
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
	if err != nil {
		return
	}
	coll = client.Database("apartments").Collection("reservations")
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

// Insert stores a record (as db.Reservation).
func Insert(request []byte) error {
	var record Reservation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := json.Unmarshal(request, &record)
	if err != nil {
		return err
	}
	_, err = coll.InsertOne(ctx, record)
	if err != nil {
		return err
	}
	return nil
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
