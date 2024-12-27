package routes

import (
	"github.com/gorilla/mux"
	"github.com/rdiego26/golang-simple-rest-api/controllers"
	"github.com/rdiego26/golang-simple-rest-api/middlewares"
	"log"
	"net/http"
)

func HandleRequests() {
	router := mux.NewRouter()

	router.Use(middlewares.ContentTypeMiddleware)

	router.HandleFunc("/", controllers.Home).Methods("GET")
	router.HandleFunc("/api/personalities", controllers.GetAllPersonalities).Methods("GET")
	router.HandleFunc("/api/personalities", controllers.CreatePersonality).Methods("POST")
	router.HandleFunc("/api/personalities/{id}", controllers.GetPersonality).Methods("GET")
	router.HandleFunc("/api/personalities/{id}", controllers.DeletePersonality).Methods("DELETE")
	router.HandleFunc("/api/personalities/{id}", controllers.EditPersonality).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", router))
}
