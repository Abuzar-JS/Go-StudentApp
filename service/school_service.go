package service

import (
	"fmt"
	"net/http"
	"student_app/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type SchoolRepository struct {
	Db *gorm.DB
}

func (s *SchoolRepository) GetSchool(context *fiber.Ctx) error {
	schoolDb := &[]models.School{}

	err := s.Db.Find(schoolDb).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"messsage": "could not get School"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "School fetched successfully",
		"data":    schoolDb,
	})
	return nil

}

func (s *SchoolRepository) GetSchoolByID(context *fiber.Ctx) error {

	id := context.Params("id")
	schoolDb := &models.School{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}
	fmt.Println("The ID is ",id)
	err := s.Db.Where("id = ?",id).First(schoolDb).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message":"could not get the school"},
		)
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message":"School Feteched Successfully",
		"data": schoolDb
	})
}
