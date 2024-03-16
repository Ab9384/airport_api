package main

import (
	"fmt"
	"net/http"
)

func main() {
	loadAirportData()
	http.HandleFunc("/autocomplete", autocompleteHandler)

	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
