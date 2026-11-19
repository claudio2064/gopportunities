package router

import (
	//"net/http"

	"github.com/claudio2064/gopportunities.git/docs"
	"github.com/claudio2064/gopportunities.git/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/claudio2064/gopportunities.git/docs"
)

func InitializeRoutes(router *gin.Engine) {

	//Initialize Handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)

	{
		v1.GET("/opening", handler.ShownOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)

	}

	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
