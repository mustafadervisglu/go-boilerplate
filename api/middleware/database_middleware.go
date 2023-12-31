package database_middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func DatabaseMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
