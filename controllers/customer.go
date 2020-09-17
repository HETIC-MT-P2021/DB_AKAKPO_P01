package controllers

import (
	"akakpo/db/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCustomer Reads all customers from database
func GetCustomer(c *gin.Context) {
	id := c.Param("id")
	response, _ := models.ReadCustomer(id)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"content": response,
	})
}
