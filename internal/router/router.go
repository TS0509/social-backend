package router

import (
	"social-backend/internal/handler"
	"social-backend/internal/middleware"
	"social-backend/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())

	// --------------------------------------------------
	// Health check
	// --------------------------------------------------
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// --------------------------------------------------
	// Services
	// --------------------------------------------------
	authService := service.NewAuthService()
	userService := service.NewUserService()

	// --------------------------------------------------
	// Handlers
	// --------------------------------------------------
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)

	// --------------------------------------------------
	// Auth Routes
	// --------------------------------------------------
	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/refresh", authHandler.Refresh)
	}

	// --------------------------------------------------
	// Protected User Routes (Auth Required)
	// --------------------------------------------------
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/profile", userHandler.Profile)
		api.PUT("/profile", userHandler.Update)
	}

	// --------------------------------------------------
	// Admin Routes (Auth + AdminOnly)
	// --------------------------------------------------
	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminOnly())
	{
		admin.GET("/stats", handler.AdminStats)
	}

	return r
}
