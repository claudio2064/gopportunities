package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"mensagem": "pong",
		})
	})
	router.Run(":8080") // escuta e serve em 0.0.0.0:8080

}
