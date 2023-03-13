package main

import (
	"jwt-authentication-golang/controller"
	"jwt-authentication-golang/database"
)

func main() {
	database.Connection()

	
	controller := controller.AuthController()
	controller.Run(":8080")
}
