package main

import (
	"data/config"
	"data/helper"
	schoolRouter "data/school/controller/router"
	studentRouter "data/student/controller/router"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Started Server! ")

	db := config.DatabaseConnection()
	validate := validator.New()

	ginRouter := gin.Default()

	schoolRouter.SchoolRouter(ginRouter, db, validate)
	studentRouter.StudentRouter(ginRouter, db, validate)

	server := &http.Server{
		Addr:    ":8080",
		Handler: ginRouter,
	}

	err := server.ListenAndServe()
	helper.ReturnError(err)

}
