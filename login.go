package main

import(
	"fmt"
	"github.com/rgarcia2304/aggreGator/internal/config"
	"errors"
)

func handlerLogin(s *state, cmd command) error{
	fmt.Println(cmd.name)
	fmt.Println(len(cmd.args))

	if len(cmd.args) < 3{
		fmt.Println("Hello")
		return errors.New("There are no users passed")
	}
	fmt.Println("blue")
	//set the new user
	err := s.cfg.SetUser(cmd.args[2])

	if err != nil{
		fmt.Println("There was an error setting the user")
		return err
	}

	//Print the new user is set
	temp, err := config.Read()
	if err != nil{
		fmt.Printf("There was an error printing the users")
		return err
	}

	s.cfg = &temp
	currUser := fmt.Sprintf("User is: %s (url is %s)", s.cfg.Username)
	fmt.Println(currUser)
	return nil
}
