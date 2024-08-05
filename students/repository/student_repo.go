package repository

import "data/students/model"

type StudentRepository interface {
	Save(student *model.Student) error
	Update(student model.Student) error
	Delete(studentId int) error
	FindById(studentId int) (Student model.Student, err error)
	FindAll() []model.Student
}
