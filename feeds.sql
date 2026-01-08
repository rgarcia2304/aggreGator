package main

import(
	"fmt"
	"errors"
	"context"
	"database/sql"
	"github.com/rgarcia2304/aggreGator/internal/database"
)

func handlerFeeds(s *state, cmd command) error{
	if len(cmd.args) != 4{
		return errors.New("Name and URL of feed must be passed")
	}

	nameArg := cmd.args[2]
	queriedName := sql.NullString{String: nameArg, Valid: true}
	ctx := context.Background()

	//check if the user exists in the database
	usr, err := s.db.GetUser(ctx, name)
	if err != nil{
		return errors.New("Issue fetching user")
	}

	//connect that user to the feed
	insertFeed, err := s.db.CreateFeedParams(ctx, database.CreateUserParams{
		Name: nameArg,
		Url: sql.NullString(String: cmd.args[3], Valid: true},
		User_id: usr.ID
	})
	fmt.Println("Feed" + insertFeed.Name)
	
	
}
