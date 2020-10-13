package main

import (
	"akakpo/db/models"
	"akakpo/db/router"

	"github.com/joho/godotenv"
)

func main() {
	// Read env vars from .env file
	env, _ := godotenv.Read(".env")

	// Attempt to connect to database - 5 tries to fallback and exit process
	models.ConnectToDatabase(env["DB_HOST"], env["DB_NAME"], env["DB_USER"], env["DB_PASSWORD"], env["DB_PORT"])

	// After router configuration, run the app
	router.Configure().Run(":" + env["API_PORT"])
}
