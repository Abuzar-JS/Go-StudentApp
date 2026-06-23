package service

import (
	schoolRepo "data/school/repository"
	"data/teacher/controller/request"
	"data/teacher/controller/response"
	"data/teacher/model"
	"data/teacher/repository"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type TeacherServiceImpl struct {
	TeacherRepository repository.TeacherRepository
	SchoolRepository  schoolRepo.SchoolRepository
	validate          *validator.Validate
}

func NewTeacherServiceImpl(teacherRepository repository.TeacherRepository, validate *validator.Validate, schoolRepository schoolRepo.SchoolRepository) TeacherService {
	return &TeacherServiceImpl{TeacherRepository: teacherRepository, validate: validate, SchoolRepository: schoolRepository}
}

func (u *TeacherServiceImpl) Create(teacher request.CreateTeacherRequest) (model.Teacher, error) {
	err := u.validate.Struct(teacher)
	if err != nil {
		return model.Teacher{}, err
	}

	_, err = u.SchoolRepository.FindById(teacher.SchoolID)
	if err != nil {
		return model.Teacher{}, fmt.Errorf("school id not found")
	}

	teacherModel := model.Teacher{
		Name:     teacher.Name,
		Subject:  teacher.Subject,
		SchoolID: teacher.SchoolID,
	}

	err = u.TeacherRepository.Save(&teacherModel)
	if err != nil {
		return model.Teacher{}, fmt.Errorf("teacher creation failed")
	}

	return teacherModel, nil
}

func (u *TeacherServiceImpl) Delete(teacherId int) error {
	err := u.TeacherRepository.Delete(teacherId)
	if err != nil {
		return fmt.Errorf("id Does not Exist")
	}
	return nil
}

func (u *TeacherServiceImpl) FindBySchoolID(schoolID int) ([]response.TeacherResponse, error) {
	_, err := u.SchoolRepository.FindById(schoolID)
	if err != nil {
		return nil, fmt.Errorf("school id not found")
	}

	result := u.TeacherRepository.FindBySchoolID(schoolID)

	var teachers []response.TeacherResponse

	for _, value := range result {
		Teacher := response.TeacherResponse{
			ID:       value.ID,
			Name:     value.Name,
			Subject:  value.Subject,
			SchoolID: value.SchoolID,
		}
		teachers = append(teachers, Teacher)
	}
	return teachers, nil
}

func (u *TeacherServiceImpl) FindById(teacherId int) (response.TeacherResponse, error) {
	Teacher, err := u.TeacherRepository.FindById(teacherId)
	if err != nil {
		return response.TeacherResponse{}, fmt.Errorf("teacher not found ")
	}
	teacherResponse := response.TeacherResponse{
		ID:       Teacher.ID,
		Name:     Teacher.Name,
		Subject:  Teacher.Subject,
		SchoolID: Teacher.SchoolID,
	}
	return teacherResponse, nil
}

func (u *TeacherServiceImpl) Update(id int, teacher request.UpdateTeacherRequest) error {
	teacherData, err := u.TeacherRepository.FindById(id)
	if err != nil {
		return fmt.Errorf("service: can't update ")
	}

	if teacher.Name != nil {
		teacherData.Name = *teacher.Name
	}

	if teacher.Subject != nil {
		teacherData.Subject = *teacher.Subject
	}

	if teacher.SchoolID != nil {
		teacherData.SchoolID = *teacher.SchoolID
	}

	if err := u.TeacherRepository.Update(id, teacherData); err != nil {
		return fmt.Errorf("update request failed: %w", err)
	}
	return nil
}
