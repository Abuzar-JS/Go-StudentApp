package router

import (
	"data/school/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(schoolController *controller.SchoolController) *gin.Engine {

	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome to School")
	})

	schoolRouter := router.Group("/api/v1")

	schoolRouter.GET("schools", schoolController.FindByAll)
	schoolRouter.GET("school/:school_id", schoolController.FindById)
	schoolRouter.POST("", schoolController.Create)
	schoolRouter.PATCH("/:school_id", schoolController.Update)
	schoolRouter.PUT("/:school_id", schoolController.Update)
	schoolRouter.DELETE("school/delete/:school_id", schoolController.Delete)

	return router
}
