package model

type School struct {
	Id   int    `gorm:"primaryKey;unique;not null" json:"id"`
	Name string `gorm:"type:varchar(255);unique;not null "`
}

func (s School) TableName() string {
	return "schools"
}
