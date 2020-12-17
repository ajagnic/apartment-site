package main

import (
	"log"
	"os"

	"github.com/ajagnic/apartment-site/server"
)

func main() {
	os.Setenv("PORT", ":8080") //temp
	port := os.Getenv("PORT")

	log.Printf("Starting server on %s", port)
	err := server.Run(port)
	if err != nil {
		log.Printf("Server error: %v", err)
	} else {
		log.Printf("Server gracefully shutdown")
	}
}
