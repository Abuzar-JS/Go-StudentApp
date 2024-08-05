package router

import (
	"data/student/controller"
	"data/student/repository"
	"data/student/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func StudentRouter(router *gin.Engine, db *gorm.DB, validate *validator.Validate) *gin.Engine {
	studentRepository := repository.NewStudentRepositoryImpl(db)

	studentService := service.NewStudentServiceImpl(studentRepository, validate)

	studentController := controller.NewStudentController(studentService)

	studentRouter := router.Group("/api/v1")

	studentRouter.GET("students", studentController.FindByAll)
	studentRouter.GET("students/:student_id", studentController.FindById)
	studentRouter.POST("/student", studentController.Create)
	studentRouter.PUT("students/:student_id", studentController.Update)
	studentRouter.DELETE("students/:student_id", studentController.Delete)

	return router
}
