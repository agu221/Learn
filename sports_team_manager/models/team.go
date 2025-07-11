package models

// Team struct represents data about a team
type Team struct {
	TeamID   int      `json:"team_id"`
	TeamName string   `json:"team_name"`
	Sport    string   `json:"sport"`
	League   string   `json:"league"`
	Division string   `json:"division"`
	Players  []Player `json:"players"`
}
