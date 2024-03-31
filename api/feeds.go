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

type FeedDupa struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Feed struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
func DBFeedsToFeedSlice(feeds []database.GetFeedsRow) []FeedDupa {
	var acc []FeedDupa

	for _, val := range feeds {
		tmpStruct := FeedDupa{
			ID:   val.ID,
			Name: val.Name,
			Url:  val.Url,
		}
		acc = append(acc, tmpStruct)
	}

	return acc
}

func (cfg *ApiConfig) CreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	decoder := json.NewDecoder(r.Body)
	params := GetFeedsPayload{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid parameters")
		return
	}

	feedPayload := database.CreateFeedParams{
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Url:       params.Url,
		UserID:    user.ID,
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), feedPayload)

	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldnt get user with provided token")
		return
	}

	respondWithJson(w, 200, DBFeedToFeed(feed))
}

func (cfg *ApiConfig) GetFeeds(w http.ResponseWriter, r *http.Request) {

	feed, err := cfg.DB.GetFeeds(r.Context())

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Someting went wrong")
		return
	}

	respondWithJson(w, 200, DBFeedsToFeedSlice(feed))
}
