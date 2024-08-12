package controller

import (
	"data/student/controller/request"
	"data/student/service"
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
	schoolID := ctx.Param("school_id")
	id, err := strconv.Atoi(schoolID)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	createStudentRequest := request.CreateStudentRequest{}
	if err := ctx.ShouldBindJSON(&createStudentRequest); err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	createStudentRequest.SchoolID = id

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

	if updateStudentRequest.Class == nil && updateStudentRequest.Name == nil && updateStudentRequest.SchoolID == nil {
		ctx.JSON(400, gin.H{
			"message": "atleast single field is required to update student",
		})
		return
	}

	err = controller.StudentService.Update(id, updateStudentRequest)
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

// Delete Controller
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
func (controller *StudentController) FindBySchoolID(ctx *gin.Context) {
	schoolID := ctx.Param("school_id")
	id, err := strconv.Atoi(schoolID)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	student, err := controller.StudentService.FindBySchoolID(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "student found",
		"data":    student,
	})

}
