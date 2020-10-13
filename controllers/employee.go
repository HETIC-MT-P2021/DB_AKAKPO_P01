package controllers

import (
	"akakpo/db/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetEmployee TODO
func GetEmployee(c *gin.Context) {
	id := c.Param("id")
	employee, readingEmployeeError := models.ReadEmployee(id)

	if readingEmployeeError != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "Cannot retreive employee data. Error message: " + readingEmployeeError.Error(),
			"data":    nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "",
			"data":    employee,
		})
	}
}
