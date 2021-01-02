package main

import (
	"log"
	"os"

	"github.com/ajagnic/apartment-site/db"
	"github.com/ajagnic/apartment-site/server"
)

func main() {
	user := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	pw := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	err := db.Connect(user, pw)
	if err != nil {
		log.Fatalf("Could not connect to database: %v\n", err)
	}

	host := os.Getenv("API_HOST")
	port := os.Getenv("API_PORT")
	addr := host + ":" + port

	log.Printf("Starting server on %s\n", addr)
	err = server.Run(addr)
	if err != nil {
		log.Printf("Server error: %v\n", err)
	}

	db.Disconnect()
}
