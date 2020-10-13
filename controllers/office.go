package controllers

import (
	"akakpo/db/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetOffice TODO
func GetOffice(c *gin.Context) {
	var err error

	id := c.Param("id")
	office, readingOfficeError := models.ReadOffice(id)

	if readingOfficeError != nil {
		err = readingOfficeError
	}

	officeEmployees, readingOfficeEmployeesError := models.ReadOfficeEmployees(office.OfficeCode)
	office.Employees = officeEmployees

	if readingOfficeEmployeesError != nil {
		err = readingOfficeEmployeesError
	}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "Cannot retreive office or its employees. Error message: " + err.Error(),
			"data":    nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "",
			"data":    office,
		})
	}
}
