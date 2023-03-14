package main

import (
	"jwt-authentication-golang/controller"
	"jwt-authentication-golang/database"
)

func main() {
	database.ConnectPostgres()

	controller := controller.Router()
	controller.Run(":8080")
}
