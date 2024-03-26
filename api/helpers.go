package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
)

type errStruct struct {
	Error string `json:"error"`
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content Type", "application/json")
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshaling JSON %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX code %d", code)
	}

	err := errStruct{
		Error: msg,
	}

	respondWithJson(w, code, err)
}

func parseHeaders(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")

	if header == "" {
		return "", errors.New("Invalid authorization token")
	}

	splitHeader := strings.Split(header, " ")

	if len(splitHeader) != 2 || splitHeader[0] != "Bearer" {
		return "", errors.New("Invalid authorization token")
	}

	return splitHeader[1], nil
}
