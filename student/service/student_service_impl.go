package service

import (
	schoolRepo "data/school/repository"
	"data/student/controller/request"
	"data/student/controller/response"
	"data/student/model"
	"data/student/repository"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type StudentServiceImpl struct {
	StudentRepository repository.StudentRepository
	validate          *validator.Validate
	SchoolRepositorty schoolRepo.SchoolRepository
}

func NewStudentServiceImpl(studentRepository repository.StudentRepository, validate *validator.Validate, schoolRepository schoolRepo.SchoolRepository) StudentService {
	return &StudentServiceImpl{
		StudentRepository: studentRepository,
		validate:          validate,
		SchoolRepositorty: schoolRepository,
	}
}

func (u *StudentServiceImpl) Create(student request.CreateStudentRequest) (model.Student, error) {
	err := u.validate.Struct(student)
	if err != nil {
		return model.Student{}, err
	}

	studentModel := model.Student{
		Name:     student.Name,
		Class:    student.Class,
		SchoolID: student.SchoolID,
	}

	err = u.StudentRepository.Save(&studentModel)
	if err != nil {
		return model.Student{}, fmt.Errorf("student creation failed")
	}

	return studentModel, nil

}

func (u *StudentServiceImpl) Delete(studentId int) error {
	err := u.StudentRepository.Delete(studentId)

	if err != nil {
		return fmt.Errorf("id Does not Exist")

	}
	return nil
}

// find all the Students in DB
func (u *StudentServiceImpl) FindBySchoolID(schoolID int) ([]response.StudentResponse, error) {

	_, err := u.SchoolRepositorty.FindById(schoolID)

	if err != nil {
		return nil, fmt.Errorf("school id not found")
	}

	result := u.StudentRepository.FindBySchoolID(schoolID)

	var students []response.StudentResponse

	for _, value := range result {
		Student := response.StudentResponse{
			ID:    value.ID,
			Name:  value.Name,
			Class: value.Class,
		}
		students = append(students, Student)
	}
	return students, nil
}

// Find Student by ID
func (u *StudentServiceImpl) FindById(studentId int) (response.StudentResponse, error) {
	Student, err := u.StudentRepository.FindById(studentId)
	if err != nil {
		return response.StudentResponse{}, fmt.Errorf("student not found ")
	}
	studentResponse := response.StudentResponse{
		ID:    Student.ID,
		Name:  Student.Name,
		Class: Student.Class,
	}
	return studentResponse, nil
}

func (u *StudentServiceImpl) Update(id int, student request.UpdateStudentRequest) error {
	studentData, err := u.StudentRepository.FindById(id)
	if err != nil {
		return fmt.Errorf("service: can't update ")
	}

	if student.Name != nil {
		studentData.Name = *student.Name
	}

	if student.Class != nil {
		studentData.Class = *student.Class
	}

	if student.SchoolID != nil {
		studentData.SchoolID = *student.SchoolID
	}
	if err := u.StudentRepository.Update(id, studentData); err != nil {
		return fmt.Errorf("update request failed: %w", err)
	}
	return nil
}
