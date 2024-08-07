package repository

import "data/student/model"

type StudentRepository interface {
	Save(student *model.Student) error
	Update(id int, student model.Student) error
	Delete(studentId int) error
	FindById(studentId int) (Student model.Student, err error)
	FindBySchoolID(studentID int) []model.Student
}
