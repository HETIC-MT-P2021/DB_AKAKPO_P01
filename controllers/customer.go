package controllers

import (
	"akakpo/db/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCustomer Reads all customers from database
func GetCustomer(c *gin.Context) {
	var err error
	id := c.Param("id")

	customer, err1 := models.ReadCustomer(id)

	if err1 != nil {
		err = err1
	}

	customerOrders, err2 := models.ReadCustomerOrders(customer.CustomerNumber)
	customer.Orders = customerOrders

	if err2 != nil {
		err = err2
	}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
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
