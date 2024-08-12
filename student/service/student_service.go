package service

import (
	"data/student/controller/request"
	"data/student/controller/response"
	"data/student/model"
)

type StudentService interface {
	Create(student request.CreateStudentRequest) (model.Student, error)
	Update(id int, student request.UpdateStudentRequest) error
	Delete(studentId int) error
	FindById(studentId int) (response.StudentResponse, error)
	FindBySchoolID(studentID int) ([]response.StudentResponse, error)
}
