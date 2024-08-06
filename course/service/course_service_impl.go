package service

import (
	"data/course/controller/request"
	"data/course/controller/response"
	"data/course/model"
	"data/course/repository"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type CourseServiceImpl struct {
	CourseRepository repository.CourseRepository
	validate         *validator.Validate
}

func NewCourseServiceImpl(courseRepository repository.CourseRepository, validate *validator.Validate) CourseService {
	return &CourseServiceImpl{CourseRepository: courseRepository, validate: validate}
}

func (u *CourseServiceImpl) Create(course request.CreateCourseRequest) (model.Course, error) {
	err := u.validate.Struct(course)
	if err != nil {
		return model.Course{}, err
	}

	courseModel := model.Course{
		Name:     course.Name,
		Class:    course.Class,
		SchoolID: course.SchoolID,
	}

	err = u.CourseRepository.Save(&courseModel)
	if err != nil {
		return model.Course{}, fmt.Errorf("course creation failed")
	}

	return courseModel, nil

}

func (u *CourseServiceImpl) Delete(courseId int) error {
	err := u.CourseRepository.Delete(courseId)

	if err != nil {
		return fmt.Errorf("id Does not Exist")

	}
	return nil
}

// find all the Courses in DB
func (u *CourseServiceImpl) FindAll() []response.CourseResponse {

	result := u.CourseRepository.FindAll()

	var courses []response.CourseResponse

	for _, value := range result {
		Course := response.CourseResponse{
			ID:    value.ID,
			Name:  value.Name,
			Class: value.Class,
		}
		courses = append(courses, Course)
	}
	return courses
}

func (u *CourseServiceImpl) FindById(courseId int) (response.CourseResponse, error) {
	Course, err := u.CourseRepository.FindById(courseId)
	if err != nil {
		return response.CourseResponse{}, fmt.Errorf("service: course not found ")
	}
	courseResponse := response.CourseResponse{
		ID:    Course.ID,
		Name:  Course.Name,
		Class: Course.Class,
	}
	return courseResponse, nil
}

func (u *CourseServiceImpl) Update(course request.UpdateCourseRequest) error {
	courseData, err := u.CourseRepository.FindById(course.ID)
	if err != nil {
		return fmt.Errorf("service: can't update ")
	}
	courseData.Name = course.Name
	u.CourseRepository.Update(courseData)
	return nil
}
