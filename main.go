package main

import(
	"github.com/rgarcia2304/aggreGator/internal/config"
	"github.com/rgarcia2304/aggreGator/internal/database"
	"fmt"
	"os"
	_ "github.com/lib/pq"
	"database/sql"	
)	

type state struct{
	cfg *config.Config 
	db *database.Queries
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

	//open connection to database
	db, err := sql.Open("postgres", stateStct.Url); 
	dbQueries := database.New(db)

	baseState := state{cfg: &stateStct, db: dbQueries}

	if err != nil{
		fmt.Println(err)
		return
	}
	
	//initialize the commands struct
	cmdsMap := make(map[string]func(*state, command) error)
	cmds := commands{validCmds: cmdsMap}

	//initialize the command struct 
	cmd := command{}
	//register the arguments to get the commands 
	cmd.args = os.Args
	fmt.Print(cmd.args)
	if len(cmd.args) < 2{
		fmt.Println("Not enough arguments passed")
		return 
	}
	cmd.name = cmd.args[1]

	err = cmds.register("login", handlerLogin)
	if err != nil{
		fmt.Println(err)
		return
	}

	err = cmds.register("register", handlerRegister)
	if err != nil{
		fmt.Println(err)
		return
	}



	err = cmds.run(&baseState, cmd)
	
	if err != nil{
		fmt.Println(err)
	}

}
