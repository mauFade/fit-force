package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Workout struct {
	ID          string
	Name        string
	Description string
	Exercises   string
	CreatedAt   time.Time
}

func NewWorkout(name string, description string, exercises string) *Workout {
	return &Workout{
		ID:          uuid.NewString(),
		Name:        name,
		Description: description,
		Exercises:   exercises,
		CreatedAt:   time.Now(),
	}
}

func (workout *Workout) Validate() error {
	if workout.ID == "" {
		return errors.New("workout ID is required")
	}

	if workout.Name == "" {
		return errors.New("workout name is required")
	}

	if workout.Description == "" {
		return errors.New("workout description is required")
	}

	if workout.Description == "" {
		return errors.New("workout requires at least one exercise")
	}

	return nil
}
