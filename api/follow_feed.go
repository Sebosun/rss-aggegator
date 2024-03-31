package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/sebosun/rss-agg/internal/database"
)

type FollowPayload struct {
	FeedId string `json:"feed_id"`
}

type FeedReturn struct {
	ID        uuid.UUID `json:"id"`
	FeedId    int32     `json:"name"`
	UserId    uuid.UUID `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func DBFollowFeed(feed database.FollowFeed) FeedReturn {
	return FeedReturn{
		ID:        feed.ID,
		FeedId:    feed.FeedID,
		UserId:    feed.UserID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
	}
}

func (cfg *ApiConfig) FollowFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	decoder := json.NewDecoder(r.Body)
	params := FollowPayload{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	feedIdInt, err := strconv.Atoi(params.FeedId)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error parsing JSON")
		return
	}

	feedInt32 := int32(feedIdInt)

	followFeedPayload := database.FollowFeedParams{
		FeedID:    feedInt32,
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	feed, err := cfg.DB.FollowFeed(r.Context(), followFeedPayload)

	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't create feed follow")
		return
	}

	respondWithJson(w, 200, DBFollowFeed(feed))
}
