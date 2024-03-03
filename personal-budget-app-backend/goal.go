package main

import (
	"time"
)

type Goal struct {
	ID         string    `json:"id"`
	UserID     string    `json:"userID"`
	Name       string    `json:"name"`
	Amount     float64   `json:"amount"`
	TargetDate time.Time `json:"targetDate"`
	Category   Category  `json:"category"`
}
