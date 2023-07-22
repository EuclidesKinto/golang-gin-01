package main

import (
	"github.com/EuclidesKinto/golang-gin-01/helper"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {

	log.Info().Msg("Started Server!")
	routes := gin.Default()
	routes.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "teste")
	})
	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
