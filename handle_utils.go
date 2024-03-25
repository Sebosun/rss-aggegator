package main

import (
	"log"
	"net/http"
)

type Status struct {
	Status string `json:"status"`
}

func handleHealthcheck(w http.ResponseWriter, r *http.Request) {
	log.Println("Received healthcheck request")
	status := Status{
		Status: "OK",
	}
	RespondWithJson(w, 200, status)
}

func handleErr(w http.ResponseWriter, r *http.Request) {
	log.Println("Received err request")
	RespondWithError(w, 404, "This endpoint should return an error")
}
