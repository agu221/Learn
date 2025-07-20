package main

import (
	"sports_team_manager/router"
	"sports_team_manager/storage"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	storage.Connect()
	r := router.SetupRouter()
	r.Run(":8080")
}
