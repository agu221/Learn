package main

import (
	"SportsNotificationApp/internal/scraper"
	"fmt"
)

func main() {
	body, err := scraper.CallSportsDBEndpoint()

	if err != nil {
		fmt.Printf("Response received: %v", body)
	}
}
