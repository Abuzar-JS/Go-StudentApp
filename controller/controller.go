package controller

import (
	"net/http"
	"student_app/service"

	"github.com/gofiber/fiber/v2"
)

type SchoolController struct {
	SchoolService service.SchoolRepository
}

func (s *SchoolController) CreateSchool(context *fiber.Ctx) error {
	type SchoolBody struct {
		Name string `json:"name"`
	}
	school := SchoolBody{}
	err := context.BodyParser(&school)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request Failed"})
		return err
	}

}
