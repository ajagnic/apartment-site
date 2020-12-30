package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func reservationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Add("Access-Control-Allow-Headers", "*")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
	} else {
		w.Header().Add("Access-Control-Allow-Origin", "*")

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("%v", err)
		}
		var jsonRequest Reservation
		err = json.Unmarshal(b, &jsonRequest)
		if err != nil {
			fmt.Printf("jsonErr:%v", err)
		}
		fmt.Printf("json: %v\n", jsonRequest)
	}
	w.Write([]byte("OK"))
}
