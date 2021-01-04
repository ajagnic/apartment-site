package main

import (
	"log"
	"os"

	"github.com/ajagnic/apartment-site/db"
	"github.com/ajagnic/apartment-site/server"
)

var (
	user   = os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	pw     = os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	dbName = os.Getenv("MONGO_INITDB_DATABASE")
	dbHost = os.Getenv("MONGO_DOMAIN_NAME")
	host   = os.Getenv("API_HOST")
	port   = os.Getenv("API_PORT")
)

func main() {
	err := db.Connect(dbHost, dbName, user, pw)
	if err != nil {
		log.Fatalf("Could not connect to database: %v\n", err)
	}

	addr := host + ":" + port

	log.Printf("Starting server on %s\n", addr)
	err = server.Run(addr)
	if err != nil {
		log.Printf("Server error: %v\n", err)
	}

	db.Disconnect()
}
