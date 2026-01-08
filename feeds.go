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

func handlerAddfeeds(s *state, cmd command) error{
	if len(cmd.args) != 4{
		return errors.New("Name and URL of feed must be passed")
	}

	queriedName := sql.NullString{String: s.cfg.Username, Valid: true}
	ctx := context.Background()

	//check if the name exists in the database
	usr, err := s.db.GetUser(ctx, queriedName)
	if err != nil{
		return errors.New("Issue fetching user")
	}

	//connect that user to the feed
	insertFeed, err := s.db.CreateFeed(ctx, database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: sql.NullString{String: cmd.args[2], Valid: true},
		Url: sql.NullString{String: cmd.args[3], Valid: true},
		UserID: usr.ID,
	})
	if err != nil{
		return err
	}
	fmt.Println("Feed" + insertFeed.Name.String)
	
	return nil
}

func handlerListFeeds(s *state, cmd command) error{
	if len(cmd.args) != 2{
		return errors.New("You can not pass arguments to the feeds command")
	}

	ctx := context.Background()
	//check if the user exists in the database
	feedLst, err := s.db.GetFeeds(ctx)
	if err != nil{
		return errors.New("Issue with executing command")
	}
	
	for _, val := range feedLst{
		name := val.Name.String
		url := val.Url.String

		//get the user who created the query 
		usr, err := s.db.GetUserByID(ctx, val.UserID)
		if err != nil{
			return errors.New("The user was not found in the database")
		}

		result := fmt.Sprintf("Name of Feed: %v , Url of Feed: %v, Name of Feed Creator: %v", name, url, usr.Name.String)
		fmt.Println(result)
	}
	return nil

}
