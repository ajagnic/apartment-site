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
	mux := http.NewServeMux()
	mux.HandleFunc("/", pingHandler)

	s := createServer(addr, mux)
	go listen(s)
	err := awaitShutdown(s)
	return err
}

func listen(s *http.Server) {
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		// Server encountered unexpected error, exit.
		log.Fatalf("server: listen: ListenAndServe: %v", err)
	}
}

func awaitShutdown(s *http.Server) error {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	// Wait for signal, then gracefully shutdown server.
	<-sigint
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := s.Shutdown(ctx)
	return err
}

func createServer(addr string, mux *http.ServeMux) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: mux,
	}
}
