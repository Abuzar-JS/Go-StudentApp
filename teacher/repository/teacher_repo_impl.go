package repository

import (
	"data/helper"
	"data/teacher/model"
	"fmt"

	"gorm.io/gorm"
)

type TeacherRepositoryImpl struct {
	Db *gorm.DB
}

func NewTeacherRepositoryImpl(Db *gorm.DB) TeacherRepository {
	return &TeacherRepositoryImpl{Db: Db}
}

func (u *TeacherRepositoryImpl) Delete(teacherId int) error {
	var teacher model.Teacher
	result := u.Db.Where("id = ?", teacherId).Delete(&teacher)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no teacher found with id %d", teacherId)
	}

	return nil
}

func (u *TeacherRepositoryImpl) FindBySchoolID(schoolID int) []model.Teacher {
	var teachers []model.Teacher
	result := u.Db.Where("school_id = ?", schoolID).Find(&teachers)
	helper.ReturnError(result.Error)
	return teachers
}

func (u *TeacherRepositoryImpl) FindById(teacherId int) (Teacher model.Teacher, err error) {
	var teacher model.Teacher
	result := u.Db.First(&teacher, teacherId)
	if result.Error != nil {
		return teacher, fmt.Errorf("teacher not found")
	}
	return teacher, nil
}

func (u *TeacherRepositoryImpl) Save(teacher *model.Teacher) error {
	result := u.Db.Create(teacher)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *TeacherRepositoryImpl) Update(id int, teacher model.Teacher) error {
	result := u.Db.Model(model.Teacher{}).Where("id=?", teacher.ID).Updates(teacher)
	if result.Error != nil {
		return fmt.Errorf("can't update")
	}

	return nil
}
