package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ReadAllCustomers Reads all customers from database
func ReadAllCustomers(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"content": "",
	})
}
