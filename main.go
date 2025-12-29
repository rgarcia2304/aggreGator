package main

import(
	"github.com/rgarcia2304/aggreGator/internal/config"
	"fmt"
)

type state struct{
	cfg *config.Config 
}

type command struct{
	name string
	args [] string
	validCmds map[string]func(*state, command) error
}

func main(){
	baseConfig, err := config.Read()
	if err != nil{
		fmt.Println(err)
	}
	
	err = baseConfig.SetUser("Rodrigo")
	if err != nil{
		fmt.Println(err)
	}

	baseConfig, err = config.Read()
	if err != nil{
		fmt.Println(err)
	}

	//print the contents of the config struct 
	currUser := fmt.Sprintf("User is: %s (url is %s)", baseConfig.Username, baseConfig.Url)
	fmt.Println(currUser)
}
