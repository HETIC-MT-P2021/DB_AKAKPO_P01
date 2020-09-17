package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ReadAllEmployees reads all the employees
func ReadAllEmployees(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Reads all the employees",
		"content": nil,
	})
}
