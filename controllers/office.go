package controllers

import (
	"akakpo/db/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetOffice Reads all offices from database
func GetOffice(c *gin.Context) {
	id := c.Param("id")
	response, _ := models.ReadOffice(id)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
"message": "",
		"data": response,
	})
}
