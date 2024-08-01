package controller

import (
	"data/helper"
	"data/school/controller/request"
	"data/school/controller/response"
	"data/school/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SchoolController struct
type SchoolController struct {
	SchoolService service.SchoolService
}

func NewSchoolController(service service.SchoolService) *SchoolController {
	return &SchoolController{
		SchoolService: service,
	}

}

// Create Controller
func (controller *SchoolController) Create(ctx *gin.Context) {
	createSchoolRequest := request.CreateSchoolRequest{}
	if err := ctx.ShouldBindJSON(&createSchoolRequest); err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	school, err := controller.SchoolService.Create(createSchoolRequest)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "School created successfully",
		"school":  school,
	})

}

// Update Controller
func (controller *SchoolController) Update(ctx *gin.Context) {
	updateSchoolRequest := request.UpdateSchoolRequest{}
	err := ctx.ShouldBindJSON(&updateSchoolRequest)
	helper.ReturnError(err)

	schoolId := ctx.Param("schoolId")
	id, err := strconv.Atoi(schoolId)
	helper.ReturnError(err)
	updateSchoolRequest.Id = id

	controller.SchoolService.Update(updateSchoolRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   updateSchoolRequest,
	}
	ctx.Header("Content-Type", "application/json")

	ctx.JSON(http.StatusOK, webResponse)
}

// Delete Controller
func (controller *SchoolController) Delete(ctx *gin.Context) {

	schoolId := ctx.Param("schoolId")
	id, err := strconv.Atoi(schoolId)
	helper.ReturnError(err)
	controller.SchoolService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   schoolId,
	}
	ctx.Header("Content-Type", "application/json")

	ctx.JSON(http.StatusOK, webResponse)
}

// FindById Controller
func (controller *SchoolController) FindById(ctx *gin.Context) {
	fmt.Println("helooooooo1")

	schoolId := ctx.Param("schoolId")
	id, err := strconv.Atoi(schoolId)
	helper.ReturnError(err)

	schoolResponse := controller.SchoolService.FindById(id)
	fmt.Println("helooooooo2", schoolResponse)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   schoolResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindByAll Controller
func (controller *SchoolController) FindByAll(ctx *gin.Context) {
	schoolResponse := controller.SchoolService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   schoolResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
