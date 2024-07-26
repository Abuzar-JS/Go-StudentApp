package main

import (
	"student_app/config"
	"student_app/models"

	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Server Started")

	db := config.ConnectionDB()

	db.Table("school").AutoMigrate(&models.School{})

}
