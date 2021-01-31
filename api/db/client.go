// Package db contains the data model and functions for interacting with a MongoDB database.
package db

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	timeout = 5 * time.Second
	retries = 10
)

var (
	client *mongo.Client
	dbm    *mongo.Database
)

var (
	host, db, table, user, pw string
)

func init() {
	host = os.Getenv("MONGO_DOMAIN_NAME")
	db = os.Getenv("MONGO_INITDB_DATABASE")
	table = os.Getenv("MONGO_APPLICATION_TABLE")
	user = os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	pw = os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
}

// Connect establishes and verifies a connection to a mongodb instance.
func Connect() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	uri := fmt.Sprintf("mongodb://%s:%s@%s:27017", user, pw, host)
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return
	}
	err = ping(ctx)
	if err != nil {
		return
	}
	dbm = client.Database(db)
	return
}

// Disconnect should be called when the client is no longer needed.
// Panics if an error occurs when calling client.Disconnect.
func Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
}

// SetBoolean updates the value of a boolean field for the specified record id.
func SetBoolean(id, field string, val bool) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	coll := dbm.Collection(table)
	filter := bson.M{"_id": objID}
	confirm := bson.D{{Key: "$set", Value: bson.D{{Key: field, Value: val}}}}
	_, err = coll.UpdateOne(ctx, filter, confirm)
	return
}

// Insert stores a record (as db.Reservation).
func Insert(request []byte) (string, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	var record Reservation
	err := json.Unmarshal(request, &record)
	if err != nil {
		return "", "", err
	}
	coll := dbm.Collection(table)
	res, err := coll.InsertOne(ctx, record)
	if err != nil {
		return "", "", err
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return id, record.Email, nil
}

// CollectDates returns a serialized map of all reserved dates per apartment.
func CollectDates() ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	coll := dbm.Collection(table)
	filter := bson.D{}
	opts := options.Find().SetProjection(bson.M{"dates": 1, "apartment": 1})
	cur, err := coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	var results []Result
	if err = cur.All(ctx, &results); err != nil {
		return nil, err
	}
	dates := sliceDates(results)
	b, err := json.Marshal(dates)
	return b, err
}

func sliceDates(recs []Result) map[string][]string {
	dateMap := make(map[string][]string)
	for _, r := range recs {
		for _, val := range r.Dates {
			dateMap[r.Apartment] = append(dateMap[r.Apartment], val)
		}
	}
	return dateMap
}

func ping(ctx context.Context) (err error) {
	for i := 0; i < retries; i++ {
		err = client.Ping(ctx, readpref.Primary())
		if err == nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
	return
}
