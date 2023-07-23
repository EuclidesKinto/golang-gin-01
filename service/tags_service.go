package service

import (
	"github.com/EuclidesKinto/golang-gin-01/data/request"
	"github.com/EuclidesKinto/golang-gin-01/data/response"
)

type TagsService interface {
	Create(tags request.CreateTagsRequest)
	Update(tags request.UpdateTagsRequest)
	Delete(tagsId int)
	FindById(tagsId int) response.TagsResponse
	FindAll() []response.TagsResponse
}
