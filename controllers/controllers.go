package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rdiego26/golang-simple-rest-api/models"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: homePage")
	fmt.Fprintln(w, "Welcome to the Home Page!")
}

func GetAllPersonalities(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getAllPersonalities")

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(models.Personalities)
	if err != nil {
		http.Error(w, "Failed to encode personalities", http.StatusInternalServerError)
	}
}

func GetPersonality(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getPersonality")

	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Println("ID:", key)

	w.Header().Set("Content-Type", "application/json")

	for _, personality := range models.Personalities {
		if personality.ID == key {
			json.NewEncoder(w).Encode(personality)
		}
	}
}
