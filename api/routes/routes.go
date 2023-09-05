package routes

import (
	"api/handlers"
	database_middleware "api/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Router(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")

	v1.Use(database_middleware.DatabaseMiddleware(db))

	v1.GET("/users", handlers.GetUsers)
	v1.POST("/createUser", handlers.CreateUser)
	v1.DELETE("/users/:id", handlers.DeleteUser)

	return router
}
