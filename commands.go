package main

import(
	"errors"
	"fmt"
)



func (c *commands) run(s *state, cmd command) error{
	fmt.Println("\n" + cmd.name + "\n")
	if cmd.name == ""{
		return errors.New("There is no function being run")
	}

	call, found := c.validCmds[cmd.name] 
	if !found{
		return errors.New("There is no function with this name")
	}
	fmt.Println(found)
	//make function call 
	err := call(s, cmd)
	if err != nil{
		return err
	}

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) error{
	if name == ""{
		return errors.New("You have passed in no name")
	}
	
	c.validCmds[name] = f
	return nil
}
