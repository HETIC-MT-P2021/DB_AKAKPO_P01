package controllers

import (
	"akakpo/db/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetEmployee Reads all employees from database
func GetEmployee(c *gin.Context) {
	id := c.Param("id")
	response, _ := models.ReadEmployee(id)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
"message": "",
		"data": response,
	})
}
