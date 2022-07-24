package main

import (
	"dc-metro-times-server/controllers"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/rail/incidents", controllers.GetRailIncidents).Methods("GET")
	router.HandleFunc("/rail/realtime", controllers.GetRailPredictions).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
	})

	handler := c.Handler(router)

	http.ListenAndServe(":5555", handler)
}
