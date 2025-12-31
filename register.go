package main

import(
	"fmt"
	"github.com/rgarcia2304/aggreGator/internal/config"
	"errors"
	"context"
	"github.com/rgarcia2304/aggreGator/internal/database"
	"github.com/google/uuid"
	"time"

)

func handlerRegister(s *state, cmd command) error{
	
	//check that there is a user passed
	if len(cmd.args) < 3{
		return errors.New("No users was passed")
	}

	//create a new user to the database 
	ctx := context.Background()
	

	//register the user 
	insertUser, err := s.db.CreateUser(ctx, database.CreateUserParams{
		id: uuid.New(),
		created_at: time.Now(),
		updated_at: time.Now(),
		name: cmd.args[2],
	})

	if err != nil{
		return err
	}
	
	fmt.Println("User was created")
	return nil
}
