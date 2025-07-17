package controllers

import (
	"errors"
	"net/http"
	"sports_team_manager/models"
	"sports_team_manager/storage"
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
	var newPlayer models.AddPlayerFormat

	// Call BindJSON to bind the received JSON to the new Team
	if err := c.BindJSON(&newPlayer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Add new team to slice
	newPlayerComplete := models.Player{PlayerID: createPlayerID(),
		PlayerName: newPlayer.PlayerName,
		DOB:        newPlayer.DOB,
		Age:        getPlayerAge(newPlayer.DOB)}
	players = append(players, newPlayerComplete)
	c.IndentedJSON(http.StatusCreated, newPlayerComplete)
}

// getTeams responds with the list of all teams as a JSON
func GetPlayers(c *gin.Context) {
	rows, err := storage.DB.Query("SELECT * FROM players")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()

	var players []models.Player
	for rows.Next() {
		var p models.Player
		if err := rows.Scan(&p.PlayerID, &p.PlayerName, &p.DOB, &p.Age); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		players = append(players, p)
	}

	c.JSON(http.StatusOK, players)
}

// getPlayerAge uses the player BOD to create their age variable
func getPlayerAge(DOB time.Time) int {
	today := time.Now()
	age := today.Year() - DOB.Year()
	if today.Month() >= DOB.Month() && today.Day() >= DOB.Day() {
		age += 1
	}
	return age
}

func createPlayerID() int {
	maxID := 0
	for _, p := range players {
		if p.PlayerID > maxID {
			maxID = p.PlayerID
		}
	}
	return maxID + 1
}

func searchPlayerInDB(playerName string) (models.Player, error) {
	for _, p := range players {
		if p.PlayerName == playerName {
			return p, nil
		}
	}
	return models.Player{}, errors.New("Player not found")
}
