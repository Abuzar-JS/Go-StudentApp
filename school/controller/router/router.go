package router

import (
	"data/school/controller"
	"data/school/repository"
	"data/school/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func SchoolRouter(router *gin.Engine, db *gorm.DB, validate *validator.Validate) *gin.Engine {
	schoolRepository := repository.NewSchoolRepositoryImpl(db)

	schoolService := service.NewSchoolServiceImpl(schoolRepository, validate)

	schoolController := controller.NewSchoolController(schoolService)

	schoolRouter := router.Group("/api/v1")

	schoolRouter.GET("/schools", schoolController.FindByAll)
	schoolRouter.GET("/schools/:school_id", schoolController.FindById)
	schoolRouter.POST("/school", schoolController.Create)
	schoolRouter.PATCH("/schools/:school_id", schoolController.Update)
	schoolRouter.PUT("/schools/:school_id", schoolController.Update)
	schoolRouter.DELETE("/schools/:school_id", schoolController.Delete)

	return router
}
