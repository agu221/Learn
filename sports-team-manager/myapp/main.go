package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/teams", getTeams)
	router.POST("/player", postPlayers)
	router.Run("localhost:8080")
}

// Team struct represents data about a team
type team struct {
	TeamID   int      `json:"team_id"`
	TeamName string   `json:"team_name"`
	Sport    string   `json:"sport"`
	League   string   `json:"league"`
	Division string   `json:"division"`
	Players  []player `json:"players"`
}

// Player struct represents data about a player
type player struct {
	PlayerID   int       `json:"player_id"`
	PlayerName string    `json:"player_name"`
	DOB        time.Time `json:"dob"`
	Age        int       `json:"int"`
}

// Player slice to seed player data
var players = []player{
	{PlayerID: 1, PlayerName: "Jason Jordan", DOB: time.Date(2001, time.August, 11, 0, 0, 0, 0, time.UTC), Age: 24},
	{PlayerID: 2, PlayerName: "Tristan Tordan", DOB: time.Date(2001, time.August, 11, 0, 0, 0, 0, time.UTC), Age: 24},
	{PlayerID: 3, PlayerName: "Dominic Dorban", DOB: time.Date(2001, time.August, 11, 0, 0, 0, 0, time.UTC), Age: 24},
}

// Teams slice to seed team date
var teams = []team{
	{TeamID: 1, TeamName: "The Birthday Boys Basketball Team", Sport: "Basketball", League: "Too Hot to Handle Deez Balls", Division: "5F", Players: players},
	{TeamID: 2, TeamName: "The Birthday Boys Soccer Team", Sport: "Soccer", League: "Its Called Soccer Not Football, Suck It Dom", Division: "1A", Players: players},
	{TeamID: 3, TeamName: "The Birthday Boys Football Team", Sport: "Football", League: "Normandeez Nuts Invitional", Division: "A", Players: players},
}

// getTeams responds with the list of all teams as a JSON
func getTeams(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, teams)
}

// postTeams adds a team from a JSON received in the request body
func postPlayers(c *gin.Context) {
	var newPlayer player

	// Call BindJSON to bind the received JSON to the new Team
	if err := c.BindJSON(&newPlayer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Add new team to slice
	players = append(players, newPlayer)
	c.IndentedJSON(http.StatusCreated, newPlayer)
}
