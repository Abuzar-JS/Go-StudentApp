package main

import (
	"data/config"
	"data/helper"
	"data/router"
	"data/school/controller"
	"data/school/model"
	"data/school/repository"
	"data/school/service"

	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Started Server! ")

	db := config.DatabaseConnection()
	validate := validator.New()

	db.AutoMigrate(&model.School{})

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
