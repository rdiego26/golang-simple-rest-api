package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rdiego26/golang-simple-rest-api/database"
	"github.com/rdiego26/golang-simple-rest-api/models"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: homePage")
	fmt.Fprintln(w, "Welcome to the Home Page!")
}

func GetAllPersonalities(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getAllPersonalities")

	var personalities []models.Personality

	w.Header().Set("Content-Type", "application/json")

	database.DB.Find(&personalities)

	err := json.NewEncoder(w).Encode(personalities)
	if err != nil {
		http.Error(w, "Failed to encode personalities", http.StatusInternalServerError)
	}
}

func GetPersonality(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getPersonality")

	vars := mux.Vars(r)
	id := vars["id"]

	w.Header().Set("Content-Type", "application/json")

	var personality models.Personality
	database.DB.Where("id = ?", id).First(&personality)

	err := json.NewEncoder(w).Encode(personality)
	if err != nil {
		http.Error(w, "Failed to encode personality", http.StatusInternalServerError)
	}
}
