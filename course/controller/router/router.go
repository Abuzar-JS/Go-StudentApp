package router

import (
	"data/course/controller"
	"data/course/repository"
	"data/course/service"
	schoolRepo "data/school/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func CourseRouter(router *gin.Engine, db *gorm.DB, validate *validator.Validate) *gin.Engine {
	courseRepository := repository.NewCourseRepositoryImpl(db)
	schoolRepository := schoolRepo.NewSchoolRepositoryImpl(db)

	courseService := service.NewCourseServiceImpl(courseRepository, validate, schoolRepository)

	courseController := controller.NewCourseController(courseService)

	courseRouter := router.Group("/api/v1/schools/:school_id/students/:student_id")
	// courseRouter := router.Group("/api/v1/schools/students/")

	courseRouter.GET("/courses", courseController.FindByStudentID)
	courseRouter.GET("/courses/:course_id", courseController.FindById)
	courseRouter.POST("/course", courseController.Create)
	courseRouter.PUT("/courses/:course_id", courseController.Update)
	courseRouter.DELETE("/courses/:course_id", courseController.Delete)

	return router
}
