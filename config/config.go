package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbName   = "coursesdb"
)

func ConnectionDB() *gorm.DB {

	dbDetails := fmt.Sprintf("host = %s port = %d user = %s password = %s dbName = %s ", host, port, user, password, dbName)
	db, err := gorm.Open(postgres.Open(dbDetails), &gorm.Config{})
	if err != nil {
		fmt.Println("Error Connecting to Database")
	}

	return db
}
