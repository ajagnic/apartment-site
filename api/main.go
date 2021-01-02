package main

import (
	"log"
	"os"

	_ "github.com/ajagnic/apartment-site/db"
	"github.com/ajagnic/apartment-site/server"
)

func main() {
	host := os.Getenv("API_HOST")
	port := os.Getenv("API_PORT")
	addr := host + ":" + port

	log.Printf("Starting server on %s", addr)
	err := server.Run(addr)
	if err != nil {
		log.Printf("Server error: %v", err)
	}
	log.Printf("Server gracefully shutdown")
}
