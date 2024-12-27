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

	if personality.ID == "" {
		http.Error(w, "Personality not found", http.StatusNotFound)
		return
	}

	err := json.NewEncoder(w).Encode(personality)
	if err != nil {
		http.Error(w, "Failed to encode personality", http.StatusInternalServerError)
	}
}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createPersonality")

	var personality models.Personality

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&personality)
	if err != nil {
		http.Error(w, "Failed to decode personality", http.StatusInternalServerError)
	}

	database.DB.Create(&personality)

	err = json.NewEncoder(w).Encode(personality)
	if err != nil {
		http.Error(w, "Failed to encode personality", http.StatusInternalServerError)
	}
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deletePersonality")

	vars := mux.Vars(r)
	id := vars["id"]

	var personality models.Personality
	database.DB.Where("id = ?", id).First(&personality)

	if personality.ID == "" {
		http.Error(w, "Personality not found", http.StatusNotFound)
		return
	}

	database.DB.Delete(&personality)

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(personality)
	if err != nil {
		http.Error(w, "Failed to encode personality", http.StatusInternalServerError)
	}
}

func EditPersonality(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: editPersonality")

	vars := mux.Vars(r)
	id := vars["id"]

	var personality models.Personality
	database.DB.Where("id = ?", id).First(&personality)

	if personality.ID == "" {
		http.Error(w, "Personality not found", http.StatusNotFound)
		return
	}

	var updatedPersonality models.Personality
	err := json.NewDecoder(r.Body).Decode(&updatedPersonality)
	if err != nil {
		http.Error(w, "Failed to decode personality", http.StatusInternalServerError)
	}

	personality.Name = updatedPersonality.Name
	personality.History = updatedPersonality.History

	database.DB.Save(&personality)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(personality)
	if err != nil {
		http.Error(w, "Failed to encode personality", http.StatusInternalServerError)
	}
}
