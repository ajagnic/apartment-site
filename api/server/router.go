package server

import (
	"net/http"
)

//Run registers route handlers and starts listening.
func Run(addr string) error {
	registerHandlers()
	s := createServer(addr)

	err := s.ListenAndServe()
	return err
}

func registerHandlers() {
	http.HandleFunc("/", index)
}

func createServer(addr string) *http.Server {
	return &http.Server{
		Addr: addr,
	}
}
