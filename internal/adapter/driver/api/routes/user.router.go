package routes

import (
	"user_microservice/internal/adapter/driver/api/handler"
	"user_microservice/internal/core/factory"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	userRepository := factory.NewUserRepository()
	userHandler := handler.NewUserHandler(userRepository)

	router.POST("/", userHandler.CreateUser)
	router.GET("/", userHandler.FindAllUsers)
	router.GET("/:id", userHandler.FindUserByID)
	router.PUT("/:id", userHandler.UpdateUser)
	router.PATCH("/:id/password", userHandler.UpdateUserPassword)
	router.DELETE("/:id", userHandler.DeleteUser)
	router.PATCH("/:id/restore", userHandler.RestoreUser)
}
