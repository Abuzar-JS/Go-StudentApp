package repository

import (
	"data/helper"
	"data/school/controller/request"
	"data/school/model"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type SchoolRepositoryImpl struct {
	Db *gorm.DB
}

func NewSchoolRepositoryImpl(Db *gorm.DB) SchoolRepository {
	return &SchoolRepositoryImpl{Db: Db}
}

func (u *SchoolRepositoryImpl) Delete(schoolId int) {
	var school model.School
	result := u.Db.Where("id = ?", schoolId).Delete(&school)

	helper.ReturnError(result.Error)

}

func (u *SchoolRepositoryImpl) FindAll() []model.School {
	var School []model.School
	result := u.Db.Find(&School)
	helper.ReturnError(result.Error)
	return School
}

func (u *SchoolRepositoryImpl) FindById(schoolId int) (School model.School, err error) {
	var school model.School
	result := u.Db.Find(&school, schoolId)
	if result != nil {
		return school, nil
	} else {
		return school, errors.New("school not found")
	}
}

func (u *SchoolRepositoryImpl) Save(school *model.School) {
	result := u.Db.Create(&school)
	fmt.Println(result.Error)
	helper.ReturnError(result.Error)
}

func (u *SchoolRepositoryImpl) Update(school model.School) {
	var updateSchool = request.UpdateSchoolRequest{
		Id:   school.Id,
		Name: school.Name,
	}

	result := u.Db.Model(&school).Updates(updateSchool)
	helper.ReturnError(result.Error)
}