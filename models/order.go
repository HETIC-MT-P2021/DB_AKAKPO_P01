package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ReadAllOrders reads all the orders
func ReadAllOrders(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Reads all the orders",
		"content": nil,
	})
}
