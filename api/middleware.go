package api

import (
	"net/http"

	"github.com/sebosun/rss-agg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *ApiConfig) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		handler(w, r, user)
	}
}
