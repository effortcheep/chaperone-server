package router

import (
	"github.com/chaperone/server/config"
	"github.com/gin-gonic/gin"
)

func SetupRouter() {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(config.ENV.SERVERPORT)
}
