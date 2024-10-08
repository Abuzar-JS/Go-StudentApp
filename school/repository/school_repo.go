package repository

import (
	"data/school/model"
)

type SchoolRepository interface {
	Save(school *model.School) error
	Update(school model.School) error
	Delete(schoolId int) error
	FindById(schoolId int) (School model.School, err error)
	FindAll() []model.School
}
