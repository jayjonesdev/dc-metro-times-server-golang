package controllers

import (
	"dc-metro-times-server/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IncidentResponse struct {
	Incidents []Incident
}

type Incident struct {
	IncidentID    string
	IncidentType  string
	Description   string
	LinesAffected string
	Min           string
}

type TrainsPredictionResponse struct {
	Trains []Train
}

type Train struct {
	Min             string
	Line            string
	Destination     string
	Car             string
	DestinationCode string
	DestinationName string
	Group           string
	LocationCode    string
	LocationName    string
}

func GetRailIncidents(c *gin.Context) {
	var incidents IncidentResponse

	WMATA_API_KEY := utils.GetEnvVar("WMATA_API_KEY")
	WMATA_API_HOST := utils.GetEnvVar("WMATA_API_HOST")

	url := fmt.Sprintf("%s/Incidents.svc/json/Incidents?api_key=%s", WMATA_API_HOST, WMATA_API_KEY)
	response, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	parseResponse(response, &incidents)
	c.IndentedJSON(http.StatusOK, incidents.Incidents)
}

func GetRailPredictions(c *gin.Context) {
	var trainPredictions TrainsPredictionResponse

	WMATA_API_KEY := utils.GetEnvVar("WMATA_API_KEY")
	WMATA_API_HOST := utils.GetEnvVar("WMATA_API_HOST")

	url := fmt.Sprintf("%s/StationPrediction.svc/json/GetPrediction/ALL?api_key=%s", WMATA_API_HOST, WMATA_API_KEY)
	response, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	parseResponse(response, &trainPredictions)

	c.IndentedJSON(http.StatusOK, trainPredictions.Trains)
}

func parseResponse(resp *http.Response, obj interface{}) {
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(bodyBytes, &obj)
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	json.Unmarshal(responseData, &obj)
}
