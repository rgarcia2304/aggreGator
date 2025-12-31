package main

import(
	"github.com/rgarcia2304/aggreGator/internal/config"
	"fmt"
	"os"
)

type state struct{
	cfg *config.Config 
}

type command struct{
	name string
	args []string
}

type commands struct{
	validCmds map[string]func(*state, command) error	
}

func main(){
	stateStct, err := config.Read()
	baseState := state{cfg: &stateStct}
	if err != nil{
		fmt.Println(err)
		return
	}
	
	//initialize the commands struct
	cmdsMap := make(map[string]func(*state, command) error)
	cmds := commands{validCmds: cmdsMap}

	//initialize the command struct 
	cmd := command{}
	
	err = cmds.register("login", handlerLogin)
	if err != nil{
		fmt.Println(err)
		return
	}

	//register the arguments to get the commands 
	cmd.args = os.Args
	
	fmt.Print(cmd.args)
	if len(cmd.args) < 2{
		fmt.Println("Not enough arguments passed")
		return 
	}
	cmd.name = cmd.args[1]
	err = cmds.run(&baseState, cmd)
	
	if err != nil{
		fmt.Println(err)
	}

}
