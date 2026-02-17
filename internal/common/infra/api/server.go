package api

import (
	"log"
	"user_microservice/internal/adapter/driver/api/routes"
	"user_microservice/internal/common/config/env"
	"user_microservice/internal/common/infra/api/middleware"
	"user_microservice/internal/common/infra/database"

	"github.com/gin-gonic/gin"
)

func Init() {
	config := env.GetConfig()

	if config.IsProduction() {
		log.Printf("Running in production mode on [%s]", config.API.URL)
		gin.SetMode(gin.ReleaseMode)
	}

	database.Connect()

	if config.Database.RunMigrations {
		database.RunMigrations()
	}

	ginRouter := gin.Default()

	ginRouter.Use(gin.Logger())
	ginRouter.Use(gin.Recovery())
	ginRouter.Use(middleware.ErrorHandlerMiddleware())

	v1Routes := ginRouter.Group("/v1")

	routes.RegisterUserRoutes(v1Routes.Group("/users"))

	ginRouter.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "healthy"})
	})

	if err := ginRouter.Run(config.API.URL); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
