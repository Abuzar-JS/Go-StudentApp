package model

import "data/school/model"

type Teacher struct {
	ID       int            `gorm:"primaryKey;unique;not null;column:id" json:"id"`
	Name     string         `gorm:"not null;column:name" json:"name"`
	Subject  string         `gorm:"not null;column:subject" json:"subject"`
	SchoolID int            `gorm:"not null;column:school_id" json:"school_id"`
	School   []model.School `gorm:"-"`
}

func (t Teacher) TableName() string {
	return "teachers"
}
