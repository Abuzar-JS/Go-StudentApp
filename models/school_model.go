package models

// School Struct
type School struct {
	Id   int    `gorm:"primaryKey;unique;not null" json:"id"`
	Name string `gorm:"not null" json:"name"`
}
