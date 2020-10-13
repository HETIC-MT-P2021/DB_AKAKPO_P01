package controllers

import (
	"akakpo/db/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCustomerAndItsOrders TODO
func GetCustomerAndItsOrders(c *gin.Context) {
	var err error

	id := c.Param("id")
	customer, readingCustomerError := models.ReadCustomer(id)

	if readingCustomerError != nil {
		err = readingCustomerError
	}

	customerOrders, readingCustomerOrdersError := models.ReadCustomerOrders(customer.CustomerNumber)
	customer.Orders = customerOrders

	if readingCustomerOrdersError != nil {
		err = readingCustomerOrdersError
	}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "Cannot retreive customer or its orders. Error message: " + err.Error(),
			"data":    nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "",
			"data":    customer,
		})
	}
}
