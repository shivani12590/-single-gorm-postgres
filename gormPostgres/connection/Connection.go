package connection

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"postgres/gormPostgres/model"
)

var dsn = "host=localhost user=postgres password=root dbname=employeeDB1 port=5432 sslmode=disable"

func Connection() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	if db != nil {
		fmt.Println("connection successful")
	}
	db.AutoMigrate(model.User{})

	return db

}
