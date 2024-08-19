package controller

import (
	"data/school/controller/request"
	"data/school/controller/response"
	"data/school/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SchoolController struct
type SchoolController struct {
	SchoolService service.SchoolService
}

// School Controller
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
			"message": err.Error(),
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

	schoolId := ctx.Param("school_id")
	id, err := strconv.Atoi(schoolId)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	updateSchoolRequest.Id = id

	err = controller.SchoolService.Update(updateSchoolRequest)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "School updated successfully",
	})
}

func (controller *SchoolController) Delete(ctx *gin.Context) {
	schoolId := ctx.Param("school_id")
	id, err := strconv.Atoi(schoolId)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = controller.SchoolService.Delete(id)
	if err != nil {
		ctx.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "school deleted successfully",
	})
}

// FindById Controller
func (controller *SchoolController) FindById(ctx *gin.Context) {

	schoolId := ctx.Param("school_id")
	id, err := strconv.Atoi(schoolId)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	school, err := controller.SchoolService.FindById(id)
	if err != nil {
		ctx.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "school found",
		"data":    school,
	})
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
