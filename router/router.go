package router

import (
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
		api.GET("/customers", healthCheck)
		// Commandes
		api.GET("/orders", healthCheck)
		// Magasins
		api.GET("/offices", healthCheck)
		// Employés des magasin
		api.GET("/employees", healthCheck)
	}

	return router
}
