package router

import (
	"social-backend/internal/handler"
	"social-backend/internal/middleware"
	"social-backend/internal/repository"
	"social-backend/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Health
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Auth DI（依赖注入）
	userRepo := &repository.UserRepository{}
	authService := &service.AuthService{UserRepo: userRepo}
	authHandler := &handler.AuthHandler{AuthService: authService}

	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	// Protected example
	r.GET("/profile", middleware.AuthMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "You are authenticated"})
	})

	return r
}
