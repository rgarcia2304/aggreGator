package main

import(
	"fmt"
	"github.com/rgarcia2304/aggreGator/internal/config"
	"errors"
	"context"
	"github.com/rgarcia2304/aggreGator/internal/database"

)

func handlerRegister(s *state, cmd command) error{
	
	//check that there is a user passed
	if len(cmd.args < 3){
		return errors.New("No users was passed")
	}

	//create a new user to the database 
	ctx := context.Background()
	
	conn, err := pgx.Connect(ctx, s.cfg.Url)
	if err != nil{
		return err
	}

	queries := 


	
}
