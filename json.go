package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content Type", "application/json")
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshaling JSON %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(data)
}

type errStruct struct {
	Error string `json:"error"`
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX code %s", code)
	}

	err := errStruct{
		Error: msg,
	}

	RespondWithJson(w, code, err)

}
