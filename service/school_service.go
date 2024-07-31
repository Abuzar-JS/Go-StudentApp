package service

import (
	"data/controller/request"
	"data/controller/response"
	"data/model"
)

type SchoolService interface {
	Create(school request.CreateSchoolRequest) model.School
	Update(school request.UpdateSchoolRequest)
	Delete(schoolId int)
	FindById(schoolId int) response.SchoolResponse
	FindAll() []response.SchoolResponse
}
