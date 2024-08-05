package router

import (
	"data/school/controller"
	"data/school/repository"
	"data/school/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func SchoolRouter(db *gorm.DB, validate *validator.Validate) *gin.Engine {

	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome to School")
	})

	schoolRepository := repository.NewSchoolRepositoryImpl(db)

	schoolService := service.NewSchoolServiceImpl(schoolRepository, validate)

	schoolController := controller.NewSchoolController(schoolService)

	schoolRouter := router.Group("/api/v1")

	schoolRouter.GET("schools", schoolController.FindByAll)
	schoolRouter.GET("school/:school_id", schoolController.FindById)
	schoolRouter.POST("", schoolController.Create)
	schoolRouter.PATCH("/:school_id", schoolController.Update)
	schoolRouter.PUT("/:school_id", schoolController.Update)
	schoolRouter.DELETE("school/delete/:school_id", schoolController.Delete)

	return router
}
