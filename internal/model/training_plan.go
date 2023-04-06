package model

import "time"

type TraningPlan struct {
	ID          string
	Name        string
	Description string
	Objective   string
	Duration    string
	CreatedAt   time.Time
}
