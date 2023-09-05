package services

import (
	"api/models"
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func CreateUser(db *gorm.DB, user *models.User) error {
	return db.Create(user).Error
}

func GetUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	fmt.Println("GetUsers", users)
	err := db.Find(&users).Error
	return users, err
}

func DeleteUser(db *gorm.DB, userID uuid.UUID) error {
	return db.Where("id = ?", userID).Delete(&models.User{}).Error
}
