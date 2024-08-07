package service

import (
	"data/course/controller/request"
	"data/course/controller/response"
	"data/course/model"
	"data/course/repository"
	schoolRepo "data/school/repository"

	"fmt"

	"github.com/go-playground/validator/v10"
)

type CourseServiceImpl struct {
	CourseRepository repository.CourseRepository
	SchoolRepository schoolRepo.SchoolRepository
	validate         *validator.Validate
}

func NewCourseServiceImpl(courseRepository repository.CourseRepository, validate *validator.Validate, schoolRepo schoolRepo.SchoolRepository) CourseService {
	return &CourseServiceImpl{
		CourseRepository: courseRepository,
		validate:         validate,
		SchoolRepository: schoolRepo,
	}
}

func (u *CourseServiceImpl) Create(course request.CreateCourseRequest) (model.Course, error) {
	err := u.validate.Struct(course)
	if err != nil {
		return model.Course{}, err
	}

	courseModel := model.Course{
		Title:     course.Title,
		StudentID: course.StudentID,
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
func (u *CourseServiceImpl) FindByStudentID(request GetCourseRequest) ([]response.CourseResponse, error) {
	_, err := u.SchoolRepository.FindById(request.SchoolID)
	if err != nil {
		return nil, fmt.Errorf("service: school ID Not Found ")
	}

	result, err := u.CourseRepository.FindByStudentID(request.StudentID)

	if err != nil {

		return nil, fmt.Errorf("service: student ID not found")
	}

	var courses []response.CourseResponse

	for _, value := range result {
		Course := response.CourseResponse{
			ID:        value.ID,
			Title:     value.Title,
			StudentID: value.StudentID,
		}
		courses = append(courses, Course)
	}
	return courses, nil
}

func (u *CourseServiceImpl) FindById(courseId int) (response.CourseResponse, error) {
	Course, err := u.CourseRepository.FindById(courseId)
	if err != nil {
		return response.CourseResponse{}, fmt.Errorf("service: course not found ")
	}
	courseResponse := response.CourseResponse{
		ID:        Course.ID,
		Title:     Course.Title,
		StudentID: Course.StudentID,
	}
	return courseResponse, nil
}

func (u *CourseServiceImpl) Update(id int, course request.UpdateCourseRequest) error {
	courseData, err := u.CourseRepository.FindById(id)
	if err != nil {
		return fmt.Errorf("service: can't update ")
	}

	if course.Title != nil {
		courseData.Title = *course.Title
	}

	if course.StudentID != nil {
		courseData.StudentID = *course.StudentID
	}
	if err := u.CourseRepository.Update(id, courseData); err != nil {
		return fmt.Errorf("update request failed: %w", err)
	}
	return nil
}
