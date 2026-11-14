package router

import "github.com/gin-gonic/gin"

func Initialize() {
	 //Inicialize Router
	router := gin.Default()
	
	//Initialize Routes
	InitializeRoutes(router)

	// Run the server
	router.Run(":8080") // escuta e serve em 0.0.0.0:8080

}
