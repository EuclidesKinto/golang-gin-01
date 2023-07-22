package repository

import "github.com/EuclidesKinto/golang-gin-01/model"

type TagsRepository interface {
	Save(tags model.Tags)
	Update(tags model.Tags)
	Delete(tagsId int)
	FindById(tagsId int) (tags model.Tags, err error)
	FinAll() []model.Tags
}
