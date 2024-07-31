package main

import (
	"data/config"
	"data/controller"
	"data/helper"
	"data/model"
	"data/repository"
	"data/router"
	"data/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Started Server! ")

	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("school").AutoMigrate(&model.School{})

	schoolRepository := repository.NewSchoolRepositoryImpl(db)

	schoolService := service.NewSchoolServiceImpl(schoolRepository, validate)

	schoolController := controller.NewSchoolController(schoolService)

	routes := router.NewRouter(schoolController)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ReturnError(err)

}
