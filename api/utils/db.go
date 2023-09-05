package utils

import (
	"api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitDB() (*gorm.DB, error) {
	dataSourceName := "host=host user=postgres password=password dbname=testGo sslmode=disable"
	db, err := gorm.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.User{})
	return db, err
}
