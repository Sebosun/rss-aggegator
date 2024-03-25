package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	godotenv.Load(".env")

	port := "8080"
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/healthz", handleHealthcheck)
	mux.HandleFunc("GET /v1/err", handleErr)

	corsMux := corsMiddleware(mux)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
