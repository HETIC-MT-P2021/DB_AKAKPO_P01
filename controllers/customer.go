package controllers

import (
	"akakpo/db/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCustomer Reads all customers from database
func GetCustomer(c *gin.Context) {
	id := c.Param("id")
	customer, _ := models.ReadCustomer(id)
	customerOrders, _ := models.ReadCustomerOrders(customer.CustomerNumber)

	customer.Orders = customerOrders

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"content": customer,
	})
}
