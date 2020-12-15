package server

import "net/http"

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Write([]byte("hello"))
	} else {
		http.Error(w, "Invalid method.", http.StatusMethodNotAllowed)
	}
}
