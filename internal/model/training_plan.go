package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type TrainingPlan struct {
	ID          string
	UserID      string
	WorkoutID   string
	Name        string
	Description string
	Objective   string
	Duration    string
	CreatedAt   time.Time
}

func NewTrainingPLan(name string, user_id string, workout_id string, description string, objective string, duration string) *TrainingPlan {
	return &TrainingPlan{
		ID:          uuid.NewString(),
		UserID:      user_id,
		WorkoutID:   workout_id,
		Name:        name,
		Description: description,
		Objective:   objective,
		Duration:    duration,
		CreatedAt:   time.Now(),
	}
}

func (training *TrainingPlan) Validate() error {
	if training.ID == "" {
		return errors.New("training plan ID is required")
	}

	if training.Name == "" {
		return errors.New("training plan name is required")
	}

	if training.Description == "" {
		return errors.New("training plan description is required")
	}

	if training.Duration == "" {
		return errors.New("training plan duration is required")
	}

	return nil
}
