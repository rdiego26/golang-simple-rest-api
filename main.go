package main

import (
	"fmt"
	"github.com/rdiego26/golang-simple-rest-api/database"
	"github.com/rdiego26/golang-simple-rest-api/routes"
)

func main() {
	database.Connect()

	fmt.Println("Initialising the application")
	routes.HandleRequests()
}
