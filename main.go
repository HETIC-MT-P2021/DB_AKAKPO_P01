package main

import (
	"akakpo/db/models"
	"akakpo/db/router"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	var host, name, user, password, port, apiPort string

	// Read env vars from env 😜
	host = os.Getenv("DB_HOST")
	name = os.Getenv("DB_NAME")
	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	port = os.Getenv("DB_PORT")
	apiPort = os.Getenv("API_PORT")

	if host == "" {
		// Read env vars from .env file instead
		env, _ := godotenv.Read(".env")

		host = env["DB_HOST"]
		name = env["DB_NAME"]
		user = env["DB_USER"]
		password = env["DB_PASSWORD"]
		port = env["DB_PORT"]
		apiPort = env["API_PORT"]
	}

	// Attempt to connect to database - 5 tries to fallback and exit process
	models.ConnectToDatabase(host, name, user, password, port)

	// After router configuration, run the app
	router.Configure().Run(":" + apiPort)
}
