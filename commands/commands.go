package main

import(
	"errors"
)
func (c *commands) run(*state, cmd commmand) error{
	if cmd.name == ""{
		return errors.New("There is no function being run")
	}

	call, err := validCmds[name] 
	if err != nil{
		return errors.New("There is no function with this name")
	}

	//make function call 
	call()
	return nil
}

func (c *commands) register(name string, f func(*state, command) error{
	if name == ""{
		return errors.New("You have passed in no name")
	}
	
	validCmds[name] = f
	return nil
}
