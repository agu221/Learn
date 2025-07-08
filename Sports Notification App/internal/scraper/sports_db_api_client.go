package scraper

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/text/date"
)

baseURL := "https://www.thesportsdb.com/api/v2/json/"

type Event struct{
	id string `json:"id"`
	idEvent string `json:"idEvent"`
	intDivision int `json:"intDivision"`
	strSport string `json:"strSport"`
	strEvent string `json:"strEvent"`
	strEventThumb string `json:"strEventThumb"`
	strTimeStamp string `json:"strTimeStamp"`
}

type Response struct {
	index int `json:"index`
	events []Event `json:"events"`
}

func GetEventsByDate(string date) (eventData) {
	
	url := baseURL + "filter/" + date
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}
	
	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	return responseObject
}
