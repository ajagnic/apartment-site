package main

import (
	"log"
	"os"

	"github.com/ajagnic/apartment-site/server"
)

func main() {
	port := os.Getenv("PORT")

	log.Printf("Starting server on %s", port)
	err := server.Run(port)
	if err != nil {
		log.Printf("Server error: %v", err)
	}
	log.Printf("Server gracefully shutdown")
}
