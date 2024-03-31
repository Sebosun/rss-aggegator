-- +goose Up
CREATE TABLE follow_feed (
    id UUID PRIMARY KEY,
    feed_id SERIAL NOT NULL REFERENCES feed(id),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE follow_feed;

