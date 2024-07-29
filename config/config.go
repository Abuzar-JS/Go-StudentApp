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
	database = "coursesdb"
)

func ConnectionDB() *gorm.DB {

	dbDetails := fmt.Sprintf("host = %s port = %d user = %s password = %s database = %s ", host, port, user, password, database)
	db, err := gorm.Open(postgres.Open(dbDetails), &gorm.Config{})
	if err != nil {
		fmt.Println("Error Connecting to Database")
	}

	return db
}
