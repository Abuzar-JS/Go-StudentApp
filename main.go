package main

import (
	"student_app/config"
	"student_app/controller"
	"student_app/models"
	"student_app/router"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Server Started")

	db := config.ConnectionDB()

	db.Table("school").AutoMigrate(&models.School{})

	app := fiber.New()
	schoolRepo := &controller.SchoolRepository{Db: db}

	routes := router.SchoolRoutes{Routes: schoolRepo}
	routes.SetupRoutes(app)

	log.Fatal().Err(app.Listen(":8080")).Msg("Server could not start")
}
