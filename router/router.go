package router

import ("github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/claudio2064/gopportunities.git/docs" 

)

func Initialize() {
	//Inicialize Router
	router := gin.Default()

	//Initialize Routes
	// ... suas rotas
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	InitializeRoutes(router)

	// Run the server
	router.Run(":8080") // escuta e serve em 0.0.0.0:8080

}
