package main

import (
	"urlshortner/database"
	"urlshortner/router"

	"github.com/joho/godotenv"  
)

func init() { //init function is a special function in Go that is called automatically before the main function is executed.
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	database.ConnectDb()
}

func main() {
	router.ClientRoutes()
}

// when the program starts, it first loads environment variables from a .env file, establishes a connection to the database, and then sets up routing for client requests. 
// This structure ensures that necessary setup tasks are performed before the server starts listening for incoming requests.
