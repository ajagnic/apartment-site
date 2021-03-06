package main

import (
	"log"

	"github.com/ajagnic/apartment-site/db"
	"github.com/ajagnic/apartment-site/email"
	"github.com/ajagnic/apartment-site/server"
)

func main() {
	err := db.Connect()
	if err != nil {
		log.Fatalf("Could not connect to database: %v\n", err)
	}

	go email.Process()

	err = server.Run()
	if err != nil {
		log.Printf("Server error: %v\n", err)
	}

	db.Disconnect()
}
