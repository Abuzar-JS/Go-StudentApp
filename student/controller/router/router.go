package router

import (
	"data/student/controller"
	"data/student/repository"
	"data/student/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func StudentRouter(db *gorm.DB, validate *validator.Validate) *gin.Engine {

	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome to Students")
	})

	studentRepository := repository.NewStudentRepositoryImpl(db)

	studentService := service.NewStudentServiceImpl(studentRepository, validate)

	studentController := controller.NewStudentController(studentService)

	studentRouter := router.Group("/api/v2")

	studentRouter.GET("students", studentController.FindByAll)
	studentRouter.GET("student/:student_id", studentController.FindById)
	studentRouter.POST("", studentController.Create)
	studentRouter.PATCH("/:student_id", studentController.Update)
	studentRouter.PUT("/:student_id", studentController.Update)
	studentRouter.DELETE("student/delete/:student_id", studentController.Delete)

	return router
}
