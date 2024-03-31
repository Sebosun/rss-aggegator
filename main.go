package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sebosun/rss-agg/api"
	"github.com/sebosun/rss-agg/internal/database"
)

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	dbURL := os.Getenv("DB_URL")

	if port == "" {
		log.Fatal("DB_URL environment variable is not set")
	}

	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	apiConfig := api.ApiConfig{
		DB: dbQueries,
	}

	mux := http.NewServeMux()
	corsMux := corsMiddleware(mux)

	mux.HandleFunc("GET /v1/healthz", apiConfig.HandleHealthcheck)
	mux.HandleFunc("GET /v1/err", apiConfig.HandleErr)

	mux.HandleFunc("GET /v1/user", apiConfig.MiddlewareAuth(apiConfig.GetUser))
	mux.HandleFunc("POST /v1/user", apiConfig.CreateUser)

	mux.HandleFunc("GET /v1/feeds", apiConfig.GetFeeds)
	mux.HandleFunc("POST /v1/feeds", apiConfig.MiddlewareAuth(apiConfig.CreateFeed))

	mux.HandleFunc("POST /v1/feed_follows", apiConfig.MiddlewareAuth(apiConfig.FollowFeed))

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
