package service

import (
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
}

func NewStudentServiceImpl(studentRepository repository.StudentRepository, validate *validator.Validate) StudentService {
	return &StudentServiceImpl{StudentRepository: studentRepository, validate: validate}
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
func (u *StudentServiceImpl) FindAll() []response.StudentResponse {

	result := u.StudentRepository.FindAll()

	var students []response.StudentResponse

	for _, value := range result {
		Student := response.StudentResponse{
			ID:    value.ID,
			Name:  value.Name,
			Class: value.Class,
		}
		students = append(students, Student)
	}
	return students
}

func (u *StudentServiceImpl) FindById(studentId int) (response.StudentResponse, error) {
	Student, err := u.StudentRepository.FindById(studentId)
	if err != nil {
		return response.StudentResponse{}, fmt.Errorf("service: student not found ")
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

	fmt.Println(student.Name)
	if student.Name != nil {
		studentData.Name = *student.Name

	}

	fmt.Println(student.Class)
	if student.Class != nil {
		studentData.Class = *student.Class
	}
	fmt.Println(student.SchoolID)
	if student.SchoolID != nil {
		studentData.SchoolID = *student.SchoolID
	}
	u.StudentRepository.Update(studentData)
	return nil
}
