package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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
		}
		var request Reservation
		err = json.Unmarshal(b, &request)
		if err != nil {
			http.Error(w, "Invalid JSON.", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusAccepted)
	}
	w.Write([]byte("OK"))
}
