package router

import (
	"data/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(schoolController *controller.SchoolController) *gin.Engine {

	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome to School")
	})

	baseRouter := router.Group("/api")
	schoolRouter := baseRouter.Group("/school")
	schoolRouter.GET("", schoolController.FindByAll)
	schoolRouter.GET("/:schoolId", schoolController.FindById)
	schoolRouter.POST("", schoolController.Create)
	schoolRouter.PATCH("/:schoolId", schoolController.Update)
	schoolRouter.PUT("/:schoolId", schoolController.Update)
	schoolRouter.DELETE("/:schoolId", schoolController.Delete)

	return router
}
