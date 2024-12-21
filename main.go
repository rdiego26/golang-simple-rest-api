package main

import (
	"fmt"
	"github.com/rdiego26/golang-simple-rest-api/models"
	"github.com/rdiego26/golang-simple-rest-api/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{ID: "1", Name: "Diego", History: "I am a software engineer"},
		{ID: "2", Name: "John", History: "I am a doctor"},
		{ID: "3", Name: "Jane", History: "I am a teacher"},
		{ID: "4", Name: "Doe", History: "I am a lawyer"},
	}

	fmt.Println("Initialising the application")
	routes.HandleRequests()
}
