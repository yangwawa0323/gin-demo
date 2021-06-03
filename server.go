package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yangwawa0323/gin-demo/controller"
	"github.com/yangwawa0323/gin-demo/middlewares"
	"github.com/yangwawa0323/gin-demo/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {

	/**
	* Same as gin.Default()
	 */
	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger())

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK!!",
		})
	})

	server.GET("/videos", func(ctx *gin.Context) {

		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8080")
}
