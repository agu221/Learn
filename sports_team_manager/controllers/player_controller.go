package controllers

import (
	"net/http"
	"sports_team_manager/models"
	"time"

	"github.com/gin-gonic/gin"
)

// Player slice to seed player data
var players = []models.Player{
	{PlayerID: 1, PlayerName: "Jason Jordan", DOB: time.Date(2001, time.August, 11, 0, 0, 0, 0, time.UTC), Age: 24},
	{PlayerID: 2, PlayerName: "Tristan Tordan", DOB: time.Date(2001, time.August, 11, 0, 0, 0, 0, time.UTC), Age: 24},
	{PlayerID: 3, PlayerName: "Dominic Dorban", DOB: time.Date(2001, time.August, 11, 0, 0, 0, 0, time.UTC), Age: 24},
}

// postPlayers adds a Player from a JSON received in the request body
func PostPlayers(c *gin.Context) {
	var newPlayer models.Player

	// Call BindJSON to bind the received JSON to the new Team
	if err := c.BindJSON(&newPlayer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Add new team to slice
	players = append(players, newPlayer)
	c.IndentedJSON(http.StatusCreated, newPlayer)
}

// getTeams responds with the list of all teams as a JSON
func GetPlayers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, players)
}
