package main

import (
	//"fmt"

	"github.com/claudio2064/gopportunities.git/config"
	"github.com/claudio2064/gopportunities.git/router"
)


var (

	logger *config.Logger

)

// @title           Gopportunities API
// @version         1.0
// @description     API for managing job openings
// @host            localhost:8080
// @BasePath        /api/v1
func main() {

	logger = config.GetLogger("main")

	//Initialize Configs
	err  := config.Init() 
	
	if err !=nil{
		logger.Errorf("config initialization error: %v", err)
		return
	}	


	//Initialize Router
	router.Initialize()

}


//https://www.youtube.com/watch?v=wyEYpX5U4Vg&t=2443s tempo parada 2:55