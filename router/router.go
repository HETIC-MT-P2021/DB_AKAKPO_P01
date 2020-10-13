package router

import (
	"akakpo/db/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Just return a simple message showing that API was successfull started
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "API is running successfully",
		"success": true,
		"data":    nil,
	})
}

// Configure configure routes and handlers and return it
func Configure() *gin.Engine {
	router := gin.New()

	api := router.Group("/v1")
	{
		// Check API status
		api.GET("/", healthCheck)

		// Customer view with its orders
		api.GET("/customers/:id", controllers.GetCustomerAndItsOrders)

		// Order view with its details
		api.GET("/orders/:id", controllers.GetOrder)

		// Employee view with its office
		api.GET("/employees/:id", controllers.GetEmployee)

		// Office view with its employees
		api.GET("/offices/:id", controllers.GetOffice)
	}

	return router
}
