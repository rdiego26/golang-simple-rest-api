package routes

import (
	"github.com/rdiego26/golang-simple-rest-api/controllers"
	"log"
	"net/http"
)

func HandleRequests() {
	router := http.NewServeMux()

	router.HandleFunc("/", controllers.Home)
	router.HandleFunc("/api/personalities", controllers.GetAllPersonalities)

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
