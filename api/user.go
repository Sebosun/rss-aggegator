package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sebosun/rss-agg/internal/database"
)

type CreateUserPayload struct {
	Name string `json:"name"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func DBUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	}
}

func (cfg *ApiConfig) CreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := CreateUserPayload{}
	err := decoder.Decode(&params)

	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		respondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldnt create user")
		return
	}

	respondWithJson(w, http.StatusOK, DBUserToUser(user))
}

func (cfg *ApiConfig) GetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := parseHeaders(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "No authorization header found")
		return
	}

	user, err := cfg.DB.GetUser(r.Context(), apiKey)

	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldnt get user with provided token")
		return
	}

	respondWithJson(w, http.StatusOK, DBUserToUser(user))
}
