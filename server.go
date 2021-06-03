package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yangwawa0323/gin-demo/controller"
	"github.com/yangwawa0323/gin-demo/middlewares"
	"github.com/yangwawa0323/gin-demo/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	logFile, err := os.Create("gin.log")
	if err != nil {
		log.Fatal("Can not create log file.")
	}

	// setup gin default writer
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
}

func main() {

	// setup log output
	setupLogOutput()

	/**
	* Same as gin.Default()
	 */
	server := gin.New()

	// add BasicAuth middleware
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK!!",
		})
	})

	server.GET("/videos", func(ctx *gin.Context) {

		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Video input is valid.",
			})
		}
	})

	server.Run(":8080")
}
