package main

import (
	"log"
	"os"

	"github.com/ajagnic/apartment-site/db"
	"github.com/ajagnic/apartment-site/server"
)

var (
	host = os.Getenv("API_HOST")
	port = os.Getenv("API_PORT")
)

func main() {
	err := db.Connect()
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
