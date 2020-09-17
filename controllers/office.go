package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ReadAllOffices Reads all offices from database
func ReadAllOffices(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"content": "",
	})
}
