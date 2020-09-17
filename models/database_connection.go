package models

import (
	"database/sql"
	"fmt"
	"time"

	// I know what I am doing
	_ "github.com/go-sql-driver/mysql"
)

// Database Represents the exported database connection obect
var Database *sql.DB

// ConnectToDatabase starts a connection with the MySQL database
func ConnectToDatabase(host string, database string, user string, password string, port string) {
	dbParameter := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database

	// Open up our database connection.
	// set up a database on local machine using phpmyadmin.
	// The database is called gomvc
	tempDB, err := sql.Open("mysql", dbParameter)

	if err != nil {
		fmt.Println("Database connection params error")
		panic(err)
	}

	err = tempDB.Ping()
	numberOfTest := 0

	for err != nil && numberOfTest < 5 {
		fmt.Println("Connection to DB did not succeed, new try")

		time.Sleep(5 * time.Second)
		tempDB, err = sql.Open("mysql", dbParameter)
		err = tempDB.Ping()

		numberOfTest++
	}

	if err != nil {
		fmt.Println("Database initialisation error")
		panic(err)
	}

	fmt.Println("Database successfully connected!")

	// defer the close till after the main function has finished
	// executing
	Database = tempDB

	// Connect and check the server version
	var version string
	Database.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)
}
