package main

import (
	"log"
	"os"

	"github.com/ajagnic/apartment-site/server"
)

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	addr := host + ":" + port

	log.Printf("Starting server on %s", addr)
	err := server.Run(addr)
	if err != nil {
		log.Printf("Server error: %v", err)
	}
	log.Printf("Server gracefully shutdown")
}
