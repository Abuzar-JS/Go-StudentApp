package router

import (
	"student_app/controller"

	"github.com/gofiber/fiber/v2"
)

type SchoolRoutes struct {
	Routes *controller.SchoolRepository
}

func (r *SchoolRoutes) SetupRoutes(app *fiber.App) {

	api := app.Group("/api")
	api.Post("/create_school", r.Routes.CreateSchool)
	api.Delete("/delete_school/:id", r.Routes.DeleteSchool)
	api.Get("/get_school/:id", r.Routes.GetSchoolByID)
	api.Get("/schools", r.Routes.GetSchool)

}
