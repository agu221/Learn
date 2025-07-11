package models

import "time"

type Player struct {
	PlayerID   int       `json:"PlayerID"`
	PlayerName string    `json:"PlayerName"`
	DOB        time.Time `json:"DOB"`
	Age        int       `json:"Age"`
}
