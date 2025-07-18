package main

import (
	"sports_team_manager/router"
	"sports_team_manager/storage"
)

func main() {

	storage.Connect()
	r := router.SetupRouter()
	r.Run(":8080")
}
