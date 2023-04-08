package usecase

import "time"

type CreateWorkoutInputDTO struct {
	ID          string
	Name        string
	Description string
	Exercises   string
}

type CreateWorkoutOutputDTO struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Exercises   string    `json:"exercises"`
	CreatedAt   time.Time `json:"created_at"`
}
