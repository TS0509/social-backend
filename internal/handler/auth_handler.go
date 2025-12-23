package handler

import (
	"net/http"
	"os"
	"time"

	"social-backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct {
	AuthService *service.AuthService
}

// =======================
//
//	REGISTER
//
// =======================
func (h *AuthHandler) Register(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email and password required"})
		return
	}

	err := h.AuthService.Register(req.Email, req.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": "email already exists"})
		return
	}

	c.JSON(200, gin.H{"message": "user registered"})
}

// =======================
//
//	LOGIN
//
// =======================
func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	if req.Email == "" || req.Password == "" {
		c.JSON(400, gin.H{"error": "email and password required"})
		return
	}

	user, err := h.AuthService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": "invalid email or password"})
		return
	}

	// =======================
	//       CREATE JWT
	// =======================
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		c.JSON(500, gin.H{"error": "JWT_SECRET not set"})
		return
	}

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		c.JSON(500, gin.H{"error": "token generation failed"})
		return
	}

	c.JSON(200, gin.H{
		"token": tokenString,
	})
}
