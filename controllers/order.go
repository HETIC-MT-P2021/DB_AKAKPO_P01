package controllers

import (
	"akakpo/db/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetOrder TODO
func GetOrder(c *gin.Context) {
	id := c.Param("id")
	orders, readingOrderError := models.ReadOrder(id)

	if readingOrderError != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "Cannot retreive order. Error message: " + readingOrderError.Error(),
			"data":    orders,
		})
	} else {
		if orders == nil {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "No order found",
				"data":    nil,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "",
				"data":    orders,
			})
		}
	}
}
