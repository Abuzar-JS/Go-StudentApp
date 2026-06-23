package controller

import (
	"data/teacher/controller/request"
	"data/teacher/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TeacherController struct {
	TeacherService service.TeacherService
}

func NewTeacherController(service service.TeacherService) *TeacherController {
	return &TeacherController{TeacherService: service}
}

func (controller *TeacherController) Create(ctx *gin.Context) {
	schoolID := ctx.Param("school_id")
	id, err := strconv.Atoi(schoolID)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	createTeacherRequest := request.CreateTeacherRequest{}
	if err := ctx.ShouldBindJSON(&createTeacherRequest); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	createTeacherRequest.SchoolID = id
	teacher, err := controller.TeacherService.Create(createTeacherRequest)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Teacher created successfully",
		"teacher": teacher,
	})
}

func (controller *TeacherController) Update(ctx *gin.Context) {
	updateTeacherRequest := request.UpdateTeacherRequest{}
	err := ctx.ShouldBindJSON(&updateTeacherRequest)

	teacherId := ctx.Param("teacher_id")
	id, err := strconv.Atoi(teacherId)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	if updateTeacherRequest.Name == nil && updateTeacherRequest.Subject == nil && updateTeacherRequest.SchoolID == nil {
		ctx.JSON(400, gin.H{"message": "atleast one field is required to update teacher"})
		return
	}

	err = controller.TeacherService.Update(id, updateTeacherRequest)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Teacher updated successfully"})
}

func (controller *TeacherController) Delete(ctx *gin.Context) {
	teacherId := ctx.Param("teacher_id")
	id, err := strconv.Atoi(teacherId)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	err = controller.TeacherService.Delete(id)
	if err != nil {
		ctx.JSON(404, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "teacher deleted successfully"})
}

func (controller *TeacherController) FindById(ctx *gin.Context) {
	teacherId := ctx.Param("teacher_id")
	id, err := strconv.Atoi(teacherId)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	teacher, err := controller.TeacherService.FindById(id)
	if err != nil {
		ctx.JSON(404, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "teacher found",
		"data":    teacher,
	})
}

func (controller *TeacherController) FindBySchoolID(ctx *gin.Context) {
	schoolID := ctx.Param("school_id")
	id, err := strconv.Atoi(schoolID)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	teacher, err := controller.TeacherService.FindBySchoolID(id)
	if err != nil {
		ctx.JSON(404, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "teacher found",
		"data":    teacher,
	})
}
