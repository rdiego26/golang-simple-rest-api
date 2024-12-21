package routes

import (
	"github.com/rdiego26/golang-simple-rest-api/controllers"
	"log"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalities", controllers.GetAllPersonalities)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
