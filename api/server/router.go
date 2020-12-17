package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

//Run registers route handlers and starts listening.
func Run(addr string) error {
	registerHandlers()
	s := createServer(addr)

	go listen(s)
	err := shutdownListener(s)
	return err
}

func registerHandlers() {
	http.HandleFunc("/", index)
}

func listen(s *http.Server) {
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("server:listen: %v", err)
	}
}

func shutdownListener(s *http.Server) error {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint
	return shutdown(s)
}

func shutdown(s *http.Server) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := s.Shutdown(ctx)
	return err
}

func createServer(addr string) *http.Server {
	return &http.Server{
		Addr: addr,
	}
}
