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
	Exercises   []string
	CreatedAt   time.Time
}

func NewWorkout(name string, description string, exercises []string) *Workout {
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
		return errors.New("ID is required")
	}

	if workout.Name == "" {
		return errors.New("Workout name is required")
	}

	if workout.Description == "" {
		return errors.New("Workout description is required")
	}

	if len(workout.Exercises) == 0 {
		return errors.New("Workout requires at least one exercise")
	}

	return nil
}
