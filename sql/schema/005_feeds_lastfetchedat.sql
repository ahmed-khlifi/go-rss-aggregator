-- +goose Up

ALTER TABLE feed_follow ADD COLUMN last_fetched_at TIMESTAMP;

-- +goose Down
ALTER TABLE feed_follow DROP COLUMN last_fetched_at;
 