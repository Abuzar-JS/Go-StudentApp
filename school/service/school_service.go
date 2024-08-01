package service

import (
	"data/school/controller/request"
	"data/school/controller/response"
	"data/school/model"
)

type SchoolService interface {
	Create(school request.CreateSchoolRequest) (model.School, error)
	Update(school request.UpdateSchoolRequest)
	Delete(schoolId int)
	FindById(schoolId int) response.SchoolResponse
	FindAll() []response.SchoolResponse
}
