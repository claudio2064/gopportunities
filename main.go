package main

import (
	"github.com/claudio2064/gopportunities.git/router"
	 
)

func main() {

	//Initialize Configs
	err  := config.Init() 
	
	if err !=nil{
		panic("Error")
	}	


	//Initialize Router
	router.Initialize()

}
