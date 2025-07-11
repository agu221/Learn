package main

import "sports_team_manager/router"

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
