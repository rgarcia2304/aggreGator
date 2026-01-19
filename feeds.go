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

func handlerAddfeeds(s *state, cmd command, user database.User) error{
	if len(cmd.args) != 4{
		return errors.New("Name and URL of feed must be passed")
	}

	ctx := context.Background()


	//connect that user to the feed
	insertFeed, err := s.db.CreateFeed(ctx, database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: sql.NullString{String: cmd.args[2], Valid: true},
		Url: sql.NullString{String: cmd.args[3], Valid: true},
		UserID: user.ID,
	})
	if err != nil{
		return err
	}

	//connect that user to the feed
	_, follow_err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: insertFeed.ID,
	})

	if follow_err != nil{
		return errors.New("Issue with following newly added feed")
	}

	fmt.Println("Feed " + insertFeed.Name.String + " created.")
	
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

func scrapeFeeds(s *state) error{
	//fetch the next feed from the database
	feed, err := s.db.GetNextFeedToFetch(context.Background()) 
	if err != nil{
		return errors.New("Could not fetch any feeds")
	}

	//mark feed as fetched
	_, err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: time.Now(),
		ID: feed.ID,
	})
	
	feedRsp, err := fetchFeed(context.Background(), feed.Url.String)

	if err != nil{
		return errors.New("Issue getting the feed response")
	}
	
	fmt.Println("Fetched: " + feedRsp.Channel.Title)
	return  nil

}
