package handler

import (
	"net/http"

	"social-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{Service: s}
}

// GET /api/profile
func (h *UserHandler) Profile(c *gin.Context) {
	userID := c.GetUint("user_id")

	user, err := h.Service.Profile(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// PUT /api/profile
func (h *UserHandler) Update(c *gin.Context) {
	userID := c.GetUint("user_id")

	var req struct {
		Avatar string `json:"avatar"`
	}
	c.BindJSON(&req)

	user, err := h.Service.Update(userID, req.Avatar)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
