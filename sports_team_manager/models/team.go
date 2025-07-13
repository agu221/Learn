package models

// Team struct represents data about a team
type Team struct {
	TeamID   int      `json:"TeamID"`
	TeamName string   `json:"TeamName"`
	Sport    string   `json:"Sport"`
	League   string   `json:"League"`
	Division string   `json:"Division"`
	Players  []Player `json:"Players"`
}

type PlayerRegistrationRequest struct {
	TeamName   string `json:"TeamName"`
	PlayerName string `json:"PlayerName"`
}
