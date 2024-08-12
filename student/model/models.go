package model

import "data/school/model"

type Student struct {
	ID       int            `gorm:"primaryKey;unique;not null;column:id" json:"id"`
	Name     string         `gorm:"not null;column:name" json:"name"`
	Class    string         `gorm:"not null;column:class" json:"class"`
	SchoolID int            `gorm:"not null;column:school_id" json:"school_id"`
	School   []model.School `gorm:"-"`
}

func (s Student) TableName() string {
	return "students"
}
