package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ReadAllCustomers reads all the customers
func ReadAllCustomers(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Reads all customers",
		"content": nil,
	})
}
