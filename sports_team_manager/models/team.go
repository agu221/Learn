package models

// Team struct represents data about a team
type Team struct {
	TeamName string `json:"TeamName"`
	Sport    string `json:"Sport"`
	League   string `json:"League"`
	Division string `json:"Division"`
	City     string `json:"City"`
}

type TeamRegisteration struct {
	TeamName      string `json:"TeamName"`
	Sport         string `json:"Sport"`
	City          string `json:"City"`
	CreatedByUser int    `json:"UserID"`
}
type PlayerRegistrationRequest struct {
	TeamName   string `json:"TeamName"`
	PlayerName string `json:"PlayerName"`
}

type TeamName struct {
	TeamName string `json:"TeamName"`
}
