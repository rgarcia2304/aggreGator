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

func handlerAddFollows(s *state, cmd command, user database.User) error{
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


	//connect that user to the feed
	insertFollow, err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil{
		return err
	}

	record := fmt.Sprintf("The name of the feed is %v, and %v has followed it", insertFollow.FeedName.String, insertFollow.UserName.String)
	fmt.Println(record)
	
	return nil
}

func handlerGetFollows(s *state, cmd command, user database.User) error{
	if len(cmd.args) != 2{
		return errors.New("No Arguments Should Be Passed")
	}

	ctx := context.Background()
	
	//grab the connected user
	usrName := user.Name

	//Get the users feeds 
	userFollows, err := s.db.GetFollowsForUser(ctx, usrName)
	if err != nil{
		return err
	}

	respPrompt := fmt.Sprintf("This is %v's feed", s.cfg.Username)
	fmt.Println(respPrompt)

	for _, val := range userFollows{
		//get the name of the feed and the name of the user s
		fmt.Println(val.FeedName.String)
	}
	return nil
}



