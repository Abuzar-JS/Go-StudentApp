package router

import (
	schoolRepo "data/school/repository"
	"data/teacher/controller"
	"data/teacher/repository"
	"data/teacher/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func TeacherRouter(router *gin.Engine, db *gorm.DB, validate *validator.Validate) *gin.Engine {
	teacherRepository := repository.NewTeacherRepositoryImpl(db)
	schoolRepository := schoolRepo.NewSchoolRepositoryImpl(db)

	teacherService := service.NewTeacherServiceImpl(teacherRepository, validate, schoolRepository)
	teacherController := controller.NewTeacherController(teacherService)

	teacherRouter := router.Group("/api/v1/schools/:school_id")

	teacherRouter.GET("/teachers", teacherController.FindBySchoolID)
	teacherRouter.GET("/teachers/:teacher_id", teacherController.FindById)
	teacherRouter.POST("/teacher", teacherController.Create)
	teacherRouter.POST("/teachers", teacherController.Create)
	teacherRouter.PUT("/teachers/:teacher_id", teacherController.Update)
	teacherRouter.DELETE("/teachers/:teacher_id", teacherController.Delete)

	return router
}
