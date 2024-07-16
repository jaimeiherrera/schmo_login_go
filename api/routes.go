package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	handlers := NewHandlers()

	// Global middlewares (applied to all routes)
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", handlers.Ping)

		// users := v1.Group("/users", middleware.Authenticate)
		// {
		// 	users.GET("/", handlers.ListUsers)
		// 	users.POST("/", handlers.CreateUser)
		// 	users.GET("/:id", handlers.GetUser)
		// 	users.PUT("/:id", handlers.UpdateUser)
		// 	users.DELETE("/:id", handlers.DeleteUser)
		// }
	}
}
