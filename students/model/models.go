package model

import "data/school/model"

type Student struct {
	Id     int            `gorm:"primaryKey;unique;not null" json:"id"`
	Name   string         `gorm:"not null" json:"name"`
	Class  string         `gorm:"not null" json:"class"`
	School []model.School `gorm:"not null" json:"school"`
}

func (s Student) TableName() string {
	return "students"
}
