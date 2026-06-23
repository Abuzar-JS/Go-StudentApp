package service

import (
	"data/teacher/controller/request"
	"data/teacher/controller/response"
	"data/teacher/model"
)

type TeacherService interface {
	Create(teacher request.CreateTeacherRequest) (model.Teacher, error)
	Update(id int, teacher request.UpdateTeacherRequest) error
	Delete(teacherId int) error
	FindById(teacherId int) (response.TeacherResponse, error)
	FindBySchoolID(schoolID int) ([]response.TeacherResponse, error)
}
