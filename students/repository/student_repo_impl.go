package repository

import (
	"data/helper"
	"data/students/controller/request"
	"data/students/model"
	"fmt"

	"gorm.io/gorm"
)

type StudentRepositoryImpl struct {
	Db *gorm.DB
}

func NewStudentRepositoryImpl(Db *gorm.DB) StudentRepository {
	return &StudentRepositoryImpl{Db: Db}
}

func (u *StudentRepositoryImpl) Delete(studentId int) error {

	var student model.Student

	result := u.Db.Where("id = ?", studentId).Delete(&student)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no student found with id %d", studentId)
	}

	return nil
}

func (u *StudentRepositoryImpl) FindAll() []model.Student {
	var Student []model.Student
	result := u.Db.Find(&Student)
	helper.ReturnError(result.Error)
	return Student
}

func (u *StudentRepositoryImpl) FindById(studentId int) (Student model.Student, err error) {
	var student model.Student
	result := u.Db.First(&student, studentId)
	if result.Error != nil {
		return student, fmt.Errorf("student not found")
	}
	return student, nil
}

func (u *StudentRepositoryImpl) Save(student *model.Student) error {
	result := u.Db.Create(student)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *StudentRepositoryImpl) Update(student model.Student) error {
	var updateStudent = request.UpdateStudentRequest{
		Id:   student.Id,
		Name: student.Name,
	}

	result := u.Db.Model(&student).Updates(updateStudent)
	if result.Error != nil {
		return fmt.Errorf("can't update")
	}
	return nil
}
