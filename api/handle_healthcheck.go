package api

import (
	"log"
	"net/http"
)

type Status struct {
	Status string `json:"status"`
}

func (cfg *ApiConfig) HandleHealthcheck(w http.ResponseWriter, _ *http.Request) {
	log.Println("Received healthcheck request")
	status := Status{
		Status: "OK",
	}
	respondWithJson(w, 200, status)
}

func (cfg *ApiConfig) HandleErr(w http.ResponseWriter, _ *http.Request) {
	log.Println("Received err request")
	respondWithError(w, 404, "This endpoint should return an error")
}
