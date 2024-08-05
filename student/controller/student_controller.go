package controller

import (
	"data/student/controller/request"
	"data/student/controller/response"
	"data/student/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// StudentController struct
type StudentController struct {
	StudentService service.StudentService
}

func NewStudentController(service service.StudentService) *StudentController {
	return &StudentController{
		StudentService: service,
	}

}

// Create Controller
func (controller *StudentController) Create(ctx *gin.Context) {
	createStudentRequest := request.CreateStudentRequest{}
	if err := ctx.ShouldBindJSON(&createStudentRequest); err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	student, err := controller.StudentService.Create(createStudentRequest)

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Student created successfully",
		"student": student,
	})

}

// Update Controller
func (controller *StudentController) Update(ctx *gin.Context) {
	updateStudentRequest := request.UpdateStudentRequest{}
	err := ctx.ShouldBindJSON(&updateStudentRequest)

	studentId := ctx.Param("student_id")
	id, err := strconv.Atoi(studentId)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	updateStudentRequest.ID = id

	err = controller.StudentService.Update(updateStudentRequest)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Student updated successfully",
	})
}

func (controller *StudentController) Delete(ctx *gin.Context) {
	studentId := ctx.Param("student_id")
	id, err := strconv.Atoi(studentId)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = controller.StudentService.Delete(id)
	if err != nil {
		ctx.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "student deleted successfully",
	})
}

// FindById Controller
func (controller *StudentController) FindById(ctx *gin.Context) {

	studentId := ctx.Param("student_id")
	id, err := strconv.Atoi(studentId)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	student, err := controller.StudentService.FindById(id)
	if err != nil {
		ctx.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "student found",
		"data":    student,
	})
}

// FindByAll Controller
func (controller *StudentController) FindByAll(ctx *gin.Context) {
	studentResponse := controller.StudentService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   studentResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
