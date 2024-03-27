package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sebosun/rss-agg/internal/database"
)

type GetFeedsPayload struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Feed struct {
	ID        int32     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
}

func DBFeedToFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
	}
}

func (cfg *ApiConfig) CreateFeed(w http.ResponseWriter, r *http.Request) {
	// TODO: middleware to handle auth
	apiKey, err := parseHeaders(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "No authorization header found")
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := GetFeedsPayload{}
	err = decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid parameters")
		return
	}

	user, err := cfg.DB.GetUser(r.Context(), apiKey)

	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldnt get user with provided token")
		return
	}

	feedPayload := database.CreateFeedParams{
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Url:       params.Url,
		UserID:    user.ID,
	}

	// TODO better query
	feed, err := cfg.DB.CreateFeed(r.Context(), feedPayload)

	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldnt get user with provided token")
		return
	}

	respondWithJson(w, 200, DBFeedToFeed(feed))
}
