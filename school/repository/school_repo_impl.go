package repository

import (
	"data/school/model"
)

type SchoolRepository interface {
	Save(school *model.School)
	Update(school model.School)
	Delete(schoolId int)
	FindById(schoolId int) (School model.School, err error)
	FindAll() []model.School
}
