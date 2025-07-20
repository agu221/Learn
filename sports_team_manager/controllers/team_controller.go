package controllers

import (
	"net/http"
	"sports_team_manager/models"
	"sports_team_manager/storage"

	"github.com/gin-gonic/gin"
)

// postPlayers adds a Player from a JSON received in the request body
func PostTeams(c *gin.Context) {
	var newTeam models.TeamRegisteration

	// Call BindJSON to bind the received JSON to the new Team
	if err := c.BindJSON(&newTeam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if storage.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No connection to database"})
		return
	}
	_, err := storage.DB.Exec("INSERT INTO teams (team_name,sport,city,created_by_user_id) VALUES($1,$2,$3,$4)", newTeam.TeamName, newTeam.Sport, newTeam.City, newTeam.CreatedByUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database insertion error"})
	}
	c.IndentedJSON(http.StatusCreated, newTeam)
}

// getTeams responds with the list of all teams as a JSON
func GetTeams(c *gin.Context) ([]models.Team, error) {
	var teams []models.Team
	if storage.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No connection to database"})
		return teams, db_connection_error()
	}

	rows, err := storage.DB.Query("SELECT team_name, sport, league, division, city FROM teams;")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return teams, err
	}

	defer rows.Close()

	for rows.Next() {
		var t models.Team
		if err := rows.Scan(&t.TeamName, &t.Sport, &t.League, &t.Division, t.City); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return teams, err
		}
		teams = append(teams, t)
	}

	return teams, nil
}

// func RegisterPlayerToTeam(c *gin.Context) {
// 	var newPlayerRegistrationRequest models.PlayerRegistrationRequest
// 	// Call BindJSON to bind the received JSON to the new Team
// 	if err := c.BindJSON(&newPlayerRegistrationRequest); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	Player, err := searchPlayerInDB(newPlayerRegistrationRequest.PlayerName)
// 	if err != nil {
// 		customErr := &CustomError{Code: 4001, Message: "Player not found in Datebase."}
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error":   true,
// 			"code":    customErr.Code,
// 			"message": customErr.Message,
// 		})
// 		return
// 	}
// 	Team, err := searchTeamInDB(newPlayerRegistrationRequest.TeamName)
// 	if err != nil {
// 		customErr := &CustomError{Code: 4001, Message: "Team not found in Datebase."}
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error":   true,
// 			"code":    customErr.Code,
// 			"message": customErr.Message,
// 		})
// 		return
// 	}
// 	Team.Players = append(Team.Players, Player)
// 	c.IndentedJSON(http.StatusCreated, newPlayerRegistrationRequest)
// }

func searchTeamInDB(c *gin.Context) (int, error) {
	var req models.TeamName
	if storage.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No connection to database"})
		return -1, db_connection_error()
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return -1, err
	}
	var TeamID int
	err := storage.DB.QueryRow("SELECT team_id FROM players WHERE LOWER(team_name) = LOWER($1);", req.TeamName).Scan(&TeamID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return -1, err
	}

	return TeamID, nil
}
