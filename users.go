package main

import(
	"fmt"
	"errors"
	"context"
)

func handlerUsers(s *state, cmd command) error{
	if len(cmd.args) != 2{
		return errors.New("You can not pass arguments to the users command")
	}

	ctx := context.Background()
	usrLst, err := s.db.GetUsers(ctx)
	if err != nil{
		return errors.New("Issue with executing command")
	}
	
	fmt.Println(" The current user is " + s.cfg.Username)
	for _, val := range usrLst{
		usr := (val.Name.String)
		if usr == s.cfg.Username{
			result := fmt.Sprintf("* " + usr + " (current)")
			fmt.Println(result)
		}else{
			result := fmt.Sprintf("* " + usr)
			fmt.Println(result)	
		}
	}
	return nil
}
