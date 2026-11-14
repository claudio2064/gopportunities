package main

import (
	"fmt"

	"github.com/claudio2064/gopportunities.git/config"
	"github.com/claudio2064/gopportunities.git/router"
)

func main() {

	//Initialize Configs
	err  := config.Init() 
	
	if err !=nil{
		fmt.Println(err)
		return
	}	


	//Initialize Router
	router.Initialize()

}
