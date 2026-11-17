package router

import (
	//"net/http"

	"github.com/claudio2064/gopportunities.git/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {

	//Initialize Handler
	handler.InitializeHandler()	


	v1 := router.Group("api/v1")

	{
		v1.GET("/opening", handler.ShownOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)

	}
}
