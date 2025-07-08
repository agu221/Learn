package scraper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseURL = "https://www.thesportsdb.com/api/v2/json/"

type Event struct {
	ID            string `json:"id"`
	IDEvent       string `json:"idEvent"`
	IntDivision   int    `json:"intDivision"`
	StrSport      string `json:"strSport"`
	StrEvent      string `json:"strEvent"`
	StrEventThumb string `json:"strEventThumb"`
	StrTimeStamp  string `json:"strTimeStamp"`
}

type Response struct {
	Index  int     `json:"index"`
	Events []Event `json:"events"`
}

func GetEventsByDate(date string) (*Response, error) {

	url := baseURL + "filter/" + date
	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("GET Failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read Failed: %w", err)
	}

	var result Response
	if err := json.Unmarshal(responseData, &result); err != nil {
		return nil, fmt.Errorf("JSON unmarshal failed: %w", err)
	}

	return &result, err
}
