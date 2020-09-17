package controllers

import (
	"akakpo/db/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetOrder Reads all orders from database
func GetOrder(c *gin.Context) {
	id := c.Param("id")
	response, _ := models.ReadOrder(id)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"content": response,
	})
}
