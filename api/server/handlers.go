package server

import (
	"io/ioutil"
	"net/http"

	"github.com/ajagnic/apartment-site/db"
)

func reservationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	if r.Method == http.MethodOptions {
		w.Header().Add("Access-Control-Allow-Headers", "*")
		w.WriteHeader(http.StatusOK)
	} else {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Could not read request body.", http.StatusBadRequest)
			return
		}
		err = db.Insert(b)
		if err != nil {
			http.Error(w, "Error saving new record.", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusAccepted)
	}
	w.Write([]byte("OK"))
}
