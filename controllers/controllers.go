package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/rdiego26/golang-simple-rest-api/models"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: homePage")
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func GetAllPersonalities(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getAllPersonalities")
	json.NewEncoder(w).Encode(models.Personalities)
}
