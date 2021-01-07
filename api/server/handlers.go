package server

import (
	"io/ioutil"
	"net/http"

	"github.com/ajagnic/apartment-site/db"
)

const (
	reservationsTable = "reservations"
)

func reservationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	if r.Method == http.MethodOptions { //~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
		w.Header().Add("Access-Control-Allow-Headers", "*")
		w.WriteHeader(http.StatusOK)

	} else if r.Method == http.MethodGet { //~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
		b, err := db.CollectDates(reservationsTable)
		if err != nil {
			http.Error(w, "Error retrieving records.", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(b)

	} else if r.Method == http.MethodPost { //~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Could not read request body.", http.StatusBadRequest)
			return
		}
		err = db.Insert(reservationsTable, b)
		if err != nil {
			http.Error(w, "Error saving new record.", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("Reservation created."))

	} else {
		http.Error(w, "Invalid method.", http.StatusMethodNotAllowed)
	}
}
