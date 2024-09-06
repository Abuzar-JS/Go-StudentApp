package model

import "data/student/model"

type Course struct {
	ID        int             `gorm:"primary_key;column:id"`
	Title     string          `gorm:"unique;not null;column:title"`
	StudentID int             `gorm:"not null;unique;column:student_id"`
	Student   []model.Student `gorm:"-"`
}

// Return Table name
func (c Course) TableName() string {
	return "courses"
}
