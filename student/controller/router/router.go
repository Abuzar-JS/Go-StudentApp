package router

import (
	schoolRepo "data/school/repository"
	"data/student/controller"
	"data/student/repository"
	"data/student/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func StudentRouter(router *gin.Engine, db *gorm.DB, validate *validator.Validate) *gin.Engine {
	studentRepository := repository.NewStudentRepositoryImpl(db)
	schoolRepository := schoolRepo.NewSchoolRepositoryImpl(db)

	studentService := service.NewStudentServiceImpl(studentRepository, validate, schoolRepository)

	studentController := controller.NewStudentController(studentService)

	studentRouter := router.Group("/api/v1/schools")

	studentRouter.GET("/:school_id/students", studentController.FindBySchoolID)
	studentRouter.GET("/:school_id/students/:student_id", studentController.FindById)
	studentRouter.POST("/:school_id/student", studentController.Create)
	studentRouter.PUT("/:school_id/students/:student_id", studentController.Update)
	studentRouter.DELETE("/:school_id/students/:student_id", studentController.Delete)

	return router
}
