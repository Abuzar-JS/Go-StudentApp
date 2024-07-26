package router

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")
	api.Post("/create_school", CreateSchool)
	api.Delete("/delete_school/:id", DeleteSchool)
	api.Get("/get_school/:id", CreateSchool)
	api.Get("/schools", GetSchool)

}
