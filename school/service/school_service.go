package service

import (
	"data/school/controller/request"
	"data/school/controller/response"
	"data/school/model"
)

type SchoolService interface {
	Create(school request.CreateSchoolRequest) (model.School, error)
	Update(school request.UpdateSchoolRequest)
	Delete(schoolId int) error
	FindById(schoolId int) (response.SchoolResponse, error)
	FindAll() []response.SchoolResponse
}
