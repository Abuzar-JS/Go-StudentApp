package controller

import (
	"data/course/controller/request"
	"data/course/controller/response"
	"data/course/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CourseController struct
type CourseController struct {
	CourseService service.CourseService
}

func NewCourseController(service service.CourseService) *CourseController {
	return &CourseController{
		CourseService: service,
	}

}

// Create Controller
func (controller *CourseController) Create(ctx *gin.Context) {
	createCourseRequest := request.CreateCourseRequest{}
	if err := ctx.ShouldBindJSON(&createCourseRequest); err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	course, err := controller.CourseService.Create(createCourseRequest)

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Course created successfully",
		"course":  course,
	})

}

// Update Controller
func (controller *CourseController) Update(ctx *gin.Context) {
	updateCourseRequest := request.UpdateCourseRequest{}
	err := ctx.ShouldBindJSON(&updateCourseRequest)

	courseId := ctx.Param("course_id")
	id, err := strconv.Atoi(courseId)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	updateCourseRequest.ID = id

	err = controller.CourseService.Update(updateCourseRequest)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Course updated successfully",
	})
}

func (controller *CourseController) Delete(ctx *gin.Context) {
	courseId := ctx.Param("course_id")
	id, err := strconv.Atoi(courseId)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = controller.CourseService.Delete(id)
	if err != nil {
		ctx.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "course deleted successfully",
	})
}

// FindById Controller
func (controller *CourseController) FindById(ctx *gin.Context) {

	courseId := ctx.Param("course_id")
	id, err := strconv.Atoi(courseId)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	course, err := controller.CourseService.FindById(id)
	if err != nil {
		ctx.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "course found",
		"data":    course,
	})
}

// FindByAll Controller
func (controller *CourseController) FindByAll(ctx *gin.Context) {
	courseResponse := controller.CourseService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   courseResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
