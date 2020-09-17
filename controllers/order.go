package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ReadAllOrders Reads all orders from database
func ReadAllOrders(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"content": "",
	})
}
