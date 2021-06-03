package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yangwawa0323/gin-demo/entity"
	"github.com/yangwawa0323/gin-demo/service"
)

/**
* The controllers is responsed to web handler function
* and manage logical including authentication, authorization,
* and delegate to access the service
 */

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
}

type controller struct {
	service service.VideoService
}

/**
* the New function is a constructor which create wrapper
* a service controller and implemented the VideoController
* interface
 */
func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

/**
* In our example, the controller has accessed to service
* directly, for the sane of simple code without too much
* controll logic.
 */
func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}
