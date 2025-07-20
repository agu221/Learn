package controllers

import (
	"net/http"
	"sports_team_manager/models"
	"sports_team_manager/storage"
	"strings"

	"github.com/gin-gonic/gin"
)

// postPlayers adds a Player from a JSON received in the request body
func PostPlayers(c *gin.Context) {
	var newPlayer models.Player

	// Call BindJSON to bind the received JSON to the new Team
	if err := c.BindJSON(&newPlayer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if storage.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No connection to database"})
		return
	}
	_, err := storage.DB.Exec("INSERT INTO players (first_name, last_name, dob) VALUES($1,$2,$3)", newPlayer.PlayerFirstName, newPlayer.PlayerLastName, newPlayer.DOB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database insertion error"})
	}
	c.IndentedJSON(http.StatusCreated, newPlayer)
}

// getTeams responds with the list of all teams as a JSON
func GetPlayers(c *gin.Context) ([]models.Player, error) {
	var players []models.Player
	if storage.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No connection to database"})
		return players, db_connection_error()
	}

	rows, err := storage.DB.Query("SELECT first_name, last_name, date_of_birth FROM players;")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return players, err
	}

	defer rows.Close()

	for rows.Next() {
		var p models.Player
		if err := rows.Scan(&p.PlayerFirstName, &p.PlayerLastName, &p.DOB); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return players, err
		}
		players = append(players, p)
	}

	return players, nil
}

func searchPlayerInDB(c *gin.Context) (int, error) {
	var req models.PlayerName
	if storage.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No connection to database"})
		return -1, db_connection_error()
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return -1, err
	}
	first_name, last_name := splitFullName(req.PlayerName)

	var playerID int
	err := storage.DB.QueryRow("SELECT player_id FROM players WHERE LOWER(first_name) = LOWER($1) AND LOWER(last_name) = LOWER($2);", first_name, last_name).Scan(&playerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return -1, err
	}

	return playerID, nil
}

func splitFullName(fullName string) (string, string) {
	parts := strings.Fields(fullName)
	if len(parts) == 0 {
		return "", ""
	} else if len(parts) == 1 {
		return parts[0], ""
	}

	// First name = first word, Last name = rest joined
	return parts[0], strings.Join(parts[1:], " ")
}
