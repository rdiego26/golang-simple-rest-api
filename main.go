package main

import (
	"fmt"
	"github.com/rdiego26/golang-simple-rest-api/models"
	"github.com/rdiego26/golang-simple-rest-api/routes"
)

func main() {
	models.Personalities = []models.Personality{
		models.Personality{ID: "1", Name: "Diego", History: "I am a software engineer"},
		models.Personality{ID: "2", Name: "John", History: "I am a doctor"},
		models.Personality{ID: "3", Name: "Jane", History: "I am a teacher"},
		models.Personality{ID: "4", Name: "Doe", History: "I am a lawyer"},
	}

	fmt.Println("Initialising the application")
	routes.HandleRequests()
}
