package service

import "github.com/yangwawa0323/gin-demo/entity"

/**
* The services is responsed to get the data or keep it,
* most situation is access to backend database.
 */

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {

	// return &videoService{}

	/*
	* the upper line will return nil to gin.JSON
	* and return a null object, so we changed code
	* by initial the videos an empty slice.
	 */

	return &videoService{videos: make([]entity.Video, 0)}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []entity.Video {
	return service.videos
}
