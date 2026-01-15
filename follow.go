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

func handlerAddFollows(s *state, cmd command) error{
	if len(cmd.args) != 3{
		return errors.New("URL of feed must be passed")
	}

	ctx := context.Background()
	
	if s.cfg.Username == ""{
		return errors.New("There is no User logged in ")
	}
	
	//grab the connected feed
	feed, err := s.db.GetFeedByURL(ctx, sql.NullString{String: cmd.args[2], Valid: true})

	if err != nil{
		return errors.New("There was an issue getting the feed at this URL")
	}

	//grab the connected user
	queriedName := sql.NullString{String: s.cfg.Username, Valid: true}
	//check if the name exists in the database
	usr, err := s.db.GetUser(ctx, queriedName)
	if err != nil{
		return errors.New("Issue fetching user")
	}

	//connect that user to the feed
	insertFollow, err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: usr.ID,
		FeedID: feed.ID,
	})

	if err != nil{
		return err
	}

	record := fmt.Sprintf("The name of the feed is %v, and %v has followed it", insertFollow.FeedName.String, insertFollow.UserName.String)
	fmt.Println(record)
	
	return nil
}
