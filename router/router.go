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

		// Fiche d'un client avec récapitulatif de ses commandes
		api.GET("/customers/:id", controllers.GetCustomerAndItsOrders)

		// Fiche d'une commande avec ses détails
		api.GET("/orders/:id", controllers.GetOrder)

		// Fiche d'un employé avec le magasin associé
		api.GET("/employees/:id", controllers.GetEmployee)

		// Fiche d'un magasin avec ses employés
		api.GET("/offices/:id", controllers.GetOffice)
	}

	return router
}
