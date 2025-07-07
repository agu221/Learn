package scraper

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func CallSportsDBEndpoint() {
	url := "https://jsonplaceholder.typicode.com/posts/1"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	fmt.Printf("Status code: %d\n", resp.StatusCode)
	fmt.Printf("Response Body: \n%s\n", body)
}
