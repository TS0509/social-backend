package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminStats(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"total_users": 123,
		"server_load": "normal",
		"uptime_days": 99,
	})
}
