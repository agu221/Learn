package controllers

import (
	"net/http"
	"sports_team_manager/models"

	"github.com/gin-gonic/gin"
)

// Teams slice to seed team date
var teams = []models.Team{
	{TeamID: 1, TeamName: "The Birthday Boys Basketball Team", Sport: "Basketball", League: "Too Hot to Handle Deez Balls", Division: "5F", Players: players},
	{TeamID: 2, TeamName: "The Birthday Boys Soccer Team", Sport: "Soccer", League: "Its Called Soccer Not Football, Suck It Dom", Division: "1A", Players: players},
	{TeamID: 3, TeamName: "The Birthday Boys Football Team", Sport: "Football", League: "Normandeez Nuts Invitional", Division: "A", Players: players},
}

// postPlayers adds a Player from a JSON received in the request body
func PostTeams(c *gin.Context) {
	var newTeam models.Team

	// Call BindJSON to bind the received JSON to the new Team
	if err := c.BindJSON(&newTeam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Add new team to slice
	teams = append(teams, newTeam)
	c.IndentedJSON(http.StatusCreated, newTeam)
}

// getTeams responds with the list of all teams as a JSON
func GetTeams(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, teams)
}
