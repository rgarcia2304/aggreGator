package main

import(
	"fmt"
	"errors"
	"context"
	"database/sql"
	"github.com/rgarcia2304/aggreGator/internal/database"
	"github.com/google/uuid"
	"time"
)

func handlerFeeds(s *state, cmd command) error{
	if len(cmd.args) != 4{
		return errors.New("Name and URL of feed must be passed")
	}

	queriedName := sql.NullString{String: cmd.args[2], Valid: true}
	ctx := context.Background()

	//check if the user exists in the database
	usr, err := s.db.GetUser(ctx, queriedName)
	if err != nil{
		return errors.New("Issue fetching user")
	}

	//connect that user to the feed
	insertFeed, err := s.db.CreateFeed(ctx, database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: queriedName,
		Url: sql.NullString{String: cmd.args[3], Valid: true},
		UserID: uuid.NullUUID{UUID: usr.ID, Valid: true},
	})
	fmt.Println("Feed" + insertFeed.Name.String)
	
	return nil
}

