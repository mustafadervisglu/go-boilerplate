package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetDB(c *gin.Context) (*gorm.DB, error) {

	dbInterface, exists := c.Get("db")

	if !exists {
		return nil, errors.New("db not found")
	}
	db, ok := dbInterface.(*gorm.DB)

	if !ok {
		return nil, errors.New("db not found")
	}
	return db, nil
}
