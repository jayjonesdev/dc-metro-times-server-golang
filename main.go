package main

import (
	"dc-metro-times-server/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func main() {
	router := gin.Default()

	router.GET("/rail/incidents", controllers.GetRailIncidents)
	router.GET("/rail/realtime", controllers.GetRailPredictions)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
	})

	http.ListenAndServe(":5555", c.Handler(router))
}
