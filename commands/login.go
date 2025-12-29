package main

import(
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error{
	if len(cmd.args) == 0{
		return errors.New("There are no arguments passed")
	}
	
	//set the new user
	err := state.cfg.SetUser(name)

	if err != nil{
		return err
	}

	//Print the new user is set
	state.cfg, err = state.cfg.Read(), err
	if err != nil{
		return err
	}

	currUser := fmt.Sprintf("User is: %s (url is %s)", state.cfg.Username)
	fmt.Println(currUser)

}
