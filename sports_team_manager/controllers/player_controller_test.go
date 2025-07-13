package controllers

import (
	"testing"
	"time"
)

func TestGetPlayerAge(t *testing.T) {
	result := getPlayerAge(time.Date(2000, time.July, 13, 0, 0, 0, 0, time.UTC))
	t.Log("Result: ", result)
}
