package repository

import (
	"data/course/controller/request"
	"data/course/model"
	"data/helper"
	"fmt"

	"gorm.io/gorm"
)

type CourseRepositoryImpl struct {
	Db *gorm.DB
}

func NewCourseRepositoryImpl(Db *gorm.DB) CourseRepository {
	return &CourseRepositoryImpl{Db: Db}
}

func (u *CourseRepositoryImpl) Delete(courseId int) error {

	var course model.Course

	result := u.Db.Where("id = ?", courseId).Delete(&course)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no course found with id %d", courseId)
	}

	return nil
}

func (u *CourseRepositoryImpl) FindAll() []model.Course {
	var Course []model.Course
	result := u.Db.Find(&Course)
	helper.ReturnError(result.Error)
	return Course
}

func (u *CourseRepositoryImpl) FindById(courseId int) (Course model.Course, err error) {
	var course model.Course
	result := u.Db.First(&course, courseId)
	if result.Error != nil {
		return course, fmt.Errorf("course not found")
	}
	return course, nil
}

func (u *CourseRepositoryImpl) Save(course *model.Course) error {
	result := u.Db.Create(course)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *CourseRepositoryImpl) Update(course model.Course) error {
	var updateCourse = request.UpdateCourseRequest{
		Name:     course.Name,
		Class:    course.Class,
		SchoolID: course.SchoolID,
	}

	result := u.Db.Model(&course).Updates(updateCourse)
	if result.Error != nil {
		return fmt.Errorf("can't update")
	}
	return nil
}
