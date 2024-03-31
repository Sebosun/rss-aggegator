-- name: CreateFeed :one
INSERT INTO feed (name, url, created_at, updated_at, user_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetFeeds :many
SELECT id, name, url FROM feed;
