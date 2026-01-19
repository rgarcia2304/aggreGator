package main

import(
	"fmt"
	"errors"
	"context"
)

func handlerReset(s *state, cmd command) error{
	if len(cmd.args) != 2{
		return errors.New("You can not pass arguments to the reset command")
	}

	ctx := context.Background()
	//check if the user exists in the database
	err := s.db.DeleteUsers(ctx)
	if err != nil{
		return errors.New("Issue with executing command")
	}
	s.cfg.SetUser("")
	fmt.Println("The database entries have been deleted")
	return nil
}
