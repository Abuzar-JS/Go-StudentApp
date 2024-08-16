package service

import (
	"data/school/controller/request"
	"data/school/controller/response"
	"data/school/model"
	"data/school/repository"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type SchoolServiceImpl struct {
	SchoolRepository repository.SchoolRepository
	validate         *validator.Validate
}

func NewSchoolServiceImpl(schoolRepository repository.SchoolRepository, validate *validator.Validate) SchoolService {
	return &SchoolServiceImpl{SchoolRepository: schoolRepository, validate: validate}
}

func (u *SchoolServiceImpl) Create(school request.CreateSchoolRequest) (model.School, error) {
	err := u.validate.Struct(school)
	if err != nil {
		return model.School{}, err
	}

	schoolModel := model.School{
		Name: school.Name,
	}

	err = u.SchoolRepository.Save(&schoolModel)
	if err != nil {
		return model.School{}, fmt.Errorf("school creation failed")
	}

	return schoolModel, nil

}

func (u *SchoolServiceImpl) Delete(schoolId int) error {
	err := u.SchoolRepository.Delete(schoolId)

	if err != nil {
		return fmt.Errorf("id Does not Exist")

	}
	return nil
}

// find all the Schools in DB
func (u *SchoolServiceImpl) FindAll() []response.SchoolResponse {

	result := u.SchoolRepository.FindAll()

	var schools []response.SchoolResponse

	for _, value := range result {
		School := response.SchoolResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		schools = append(schools, School)
	}
	return schools
}

func (u *SchoolServiceImpl) FindById(schoolId int) (response.SchoolResponse, error) {
	School, err := u.SchoolRepository.FindById(schoolId)
	if err != nil {
		return response.SchoolResponse{}, fmt.Errorf("service: school not found ")
	}
	schoolResponse := response.SchoolResponse{
		Id:   School.Id,
		Name: School.Name,
	}
	return schoolResponse, nil
}

// Update School
func (u *SchoolServiceImpl) Update(school request.UpdateSchoolRequest) error {
	schoolData, err := u.SchoolRepository.FindById(school.Id)
	if err != nil {
		return fmt.Errorf("service: can't update ")
	}
	schoolData.Name = school.Name
	u.SchoolRepository.Update(schoolData)
	return nil
}
