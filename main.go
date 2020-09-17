package main

import (
	"akakpo/db/models"
	"akakpo/db/router"

	"github.com/joho/godotenv"
)

func main() {
	env, _ := godotenv.Read(".env")

	models.ConnectToDatabase(env["DB_HOST"], env["DB_NAME"], env["DB_USER"], env["DB_PASSWORD"], env["DB_PORT"])

	router.Configure().Run(":" + env["API_PORT"])
}
