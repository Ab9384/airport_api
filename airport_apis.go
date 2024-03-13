package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

type Airport struct {
	ICAO      string  `json:"icao"`
	IATA      string  `json:"iata"`
	Name      string  `json:"name"`
	City      string  `json:"city"`
	State     string  `json:"state"`
	Country   string  `json:"country"`
	Elevation int     `json:"elevation"`
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lon"`
	TZ        string  `json:"tz"`
}

var airportList []Airport

func loadAirportData() {
	file, err := os.Open("./airport_data.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&airportList)
	if err != nil {
		panic(err)
	}

}

func searchAirports(query string) []Airport {
	query = strings.ToLower(query)
	var results = []Airport{}
	query = strings.ToLower(query)

	for _, airport := range airportList {
		if strings.Contains(strings.ToLower(airport.ICAO), query) ||
			strings.Contains(strings.ToLower(airport.IATA), query) ||
			strings.Contains(strings.ToLower(airport.Name), query) ||
			strings.Contains(strings.ToLower(airport.City), query) ||
			strings.Contains(strings.ToLower(airport.State), query) ||
			strings.Contains(strings.ToLower(airport.Country), query) {
			results = append(results, airport)
		}
	}

	return results
}

func autocompleteHandler(responseWriter http.ResponseWriter, request *http.Request) {
	// get the search query from the URL
	query := request.URL.Query().Get("search")
	responseWriter.Header().Set("Content-Type", "application/json")

	if query == "" {
		responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Write([]byte(`{"error": "missing query parameter"}`))
		return
	}

	results := searchAirports(query)
	jsonResult, err := json.Marshal(results)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	responseWriter.Write(jsonResult)
}
