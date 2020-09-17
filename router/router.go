package router

import (
	"akakpo/db/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "API is running successfully",
		"success": true,
	})
}

// Configure configure routes and handlers and return it
func Configure() *gin.Engine {
	router := gin.New()

	api := router.Group("/v1")
	{
		// Check API status
		api.GET("/", healthCheck)

		// Dashboard routes
		// Clients
		api.GET("/customers/:id", controllers.GetCustomer)
		// Commandes
		api.GET("/orders/:id", controllers.GetOrder)
		// Magasins
		api.GET("/offices/:id", controllers.GetOffice)
		// Employés des magasin
		api.GET("/employees/:id", controllers.GetEmployee)
	}

	return router
}
