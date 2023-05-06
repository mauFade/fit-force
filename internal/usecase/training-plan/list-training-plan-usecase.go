package usecase

import "time"

type ListTrainingPLanOutputDTO struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	WorkoutID   string    `json:"workout_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Objective   string    `json:"objective"`
	Duration    string    `json:"duration"`
	CreatedAt   time.Time `json:"created_at"`
}
