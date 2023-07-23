package main

import (
	"github.com/EuclidesKinto/golang-gin-01/config"
	"github.com/EuclidesKinto/golang-gin-01/controller"
	"github.com/EuclidesKinto/golang-gin-01/helper"
	"github.com/EuclidesKinto/golang-gin-01/model"
	"github.com/EuclidesKinto/golang-gin-01/repository"
	"github.com/EuclidesKinto/golang-gin-01/router"
	"github.com/EuclidesKinto/golang-gin-01/service"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {

	log.Info().Msg("Started Server!")
	//Databse
	db := config.DataBaseConnection()
	validate := validator.New()
	db.Table("tags").AutoMigrate(&model.Tags{})

	// Repository
	tagsRepository := repository.NewTagsRepositoryImpl(db)

	// Service
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)

	// Controller
	tagsController := controller.NewTagsController(tagsService)

	// Router
	routes := router.NewRouter(tagsController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
