package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ReadAllOffices reads all the offices
func ReadAllOffices(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Reads all the offices",
		"content": nil,
	})
}
