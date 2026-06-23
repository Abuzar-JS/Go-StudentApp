package repository

import "data/teacher/model"

type TeacherRepository interface {
	Save(teacher *model.Teacher) error
	Update(id int, teacher model.Teacher) error
	Delete(teacherId int) error
	FindById(teacherId int) (Teacher model.Teacher, err error)
	FindBySchoolID(schoolID int) []model.Teacher
}
