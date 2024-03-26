package api

import (
	"github.com/sebosun/rss-agg/internal/database"
)

type ApiConfig struct {
	DB *database.Queries
}
