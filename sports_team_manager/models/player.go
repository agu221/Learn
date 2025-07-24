package models

import "time"

type Player struct {
	PlayerFirstName string    `json:"PlayerFirstName"`
	PlayerLastName  string    `json:"PlayerLastName"`
	DOB             time.Time `json:"DOB"`
}

type PlayerName struct {
	PlayerName string `json:"PlayerName"`
}

type PlayerTeams struct {
	PlayerID           int                              `json:"PlayerID"`
	RegisteredTeamsArr []RegisteredTeamsPlayerDashboard `json:"RegisteredTeams"`
}
