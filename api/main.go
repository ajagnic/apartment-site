package main

import (
	"log"
	"os"

	"github.com/ajagnic/apartment-site/db"
	"github.com/ajagnic/apartment-site/email"
	"github.com/ajagnic/apartment-site/server"
)

var (
	host        = os.Getenv("API_HOST")
	port        = os.Getenv("API_PORT")
	emailHost   = os.Getenv("EMAIL_HOST")
	emailPort   = os.Getenv("EMAIL_PORT")
	emailSender = os.Getenv("EMAIL_SENDER_ADDR")
	emailPW     = os.Getenv("EMAIL_SENDER_PW")
)

func main() {
	err := db.Connect()
	if err != nil {
		log.Fatalf("Could not connect to database: %v\n", err)
	}

	email.Initialize(emailHost, emailPort, emailSender, emailPW)

	addr := host + ":" + port

	log.Printf("Starting server on %s\n", addr)
	err = server.Run(addr)
	if err != nil {
		log.Printf("Server error: %v\n", err)
	}

	db.Disconnect()
}
