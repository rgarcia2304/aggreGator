-- +goose Up 
CREATE TABLE feeds(
	name TEXT,
	url TEXT UNIQUE, 
	user_id UUID REFERENCES users (id) ON DELETE CASCADE
);
