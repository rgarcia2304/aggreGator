package main

import(
	"fmt"
	"errors"
	"context"
	"github.com/rgarcia2304/aggreGator/internal/database"
	"github.com/google/uuid"
	"time"
	"database/sql"

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
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: sql.NullString{String: cmd.args[2], Valid: true},
	})

	if err != nil{
		return err
	}
	
	fmt.Println(insertUser)
	fmt.Println("User was created")
	return nil
}
