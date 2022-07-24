package controllers

import (
	"dc-metro-times-server/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

func GetRailIncidents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var incidents IncidentResponse

	WMATA_API_KEY := utils.GetEnvVar("WMATA_API_KEY")
	WMATA_API_HOST := utils.GetEnvVar("WMATA_API_HOST")

	url := fmt.Sprintf("%s/Incidents.svc/json/Incidents?api_key=%s", WMATA_API_HOST, WMATA_API_KEY)
	response, httpErr := http.Get(url)

	if httpErr != nil {
		log.Fatalln(httpErr)
	}

	parseResponse(response, &incidents)

	jsonResp, jsonErr := json.Marshal(incidents.Incidents)

	if jsonErr != nil {
		log.Fatalln(jsonErr)
	}

	w.Write(jsonResp)
}

func GetRailPredictions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var trainPredictions TrainsPredictionResponse

	WMATA_API_KEY := utils.GetEnvVar("WMATA_API_KEY")
	WMATA_API_HOST := utils.GetEnvVar("WMATA_API_HOST")

	url := fmt.Sprintf("%s/StationPrediction.svc/json/GetPrediction/ALL?api_key=%s", WMATA_API_HOST, WMATA_API_KEY)
	response, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	parseResponse(response, &trainPredictions)

	jsonResp, jsonErr := json.Marshal(trainPredictions.Trains)

	if jsonErr != nil {
		log.Fatalln(jsonErr)
	}

	w.Write(jsonResp)
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
