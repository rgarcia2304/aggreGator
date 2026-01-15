-- +goose Up
CREATE TABLE feed_follows(
	id UUID PRIMARY KEY, 
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	user_id UUID NOT NULL, 
	feed_id UUID NOT NULL, 
	UNIQUE(user_id, feed_id),
	Foreign KEY (user_id)
	REFERENCES users(id) ON DELETE CASCADE,
	Foreign KEY (feed_id)
	REFERENCES feeds(id) ON DELETE CASCADE
);
