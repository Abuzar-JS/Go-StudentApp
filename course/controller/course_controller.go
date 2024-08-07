package controller

import (
	"data/course/controller/request"
	"data/course/service"
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
	if updateCourseRequest.Title == nil && updateCourseRequest.StudentID == nil {
		ctx.JSON(400, gin.H{
			"message": "atleast one field is required to update course",
		})
		return
	}

	err = controller.CourseService.Update(id, updateCourseRequest)
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
func (controller *CourseController) FindByStudentID(ctx *gin.Context) {
	courseID := ctx.Param("student_id")
	id, err := strconv.Atoi(courseID)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	course := controller.CourseService.FindByStudentID(id)

	ctx.JSON(200, gin.H{
		"message": "course found",
		"data":    course,
	})

}
