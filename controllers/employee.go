package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ReadAllEmployees Reads all employees from database
func ReadAllEmployees(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"content": "",
	})
}
