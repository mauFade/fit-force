package usecase

import (
	"time"

	"github.com/mauFade/fit-force/internal/infra/repository"
)

type ListWorkoutOutputDTO struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Exercises   string    `json:"exercises"`
	CreatedAt   time.Time `json:"created_at"`
}

type ListWorkoutsUseCase struct {
	WorkoutRepository repository.WorkoutDatabaseRepository
}

func NewListWorkoutsUseCase(repo repository.WorkoutDatabaseRepository) *ListWorkoutsUseCase {
	return &ListWorkoutsUseCase{
		WorkoutRepository: repo,
	}
}

func (use_case *ListWorkoutsUseCase) Execute() []*ListWorkoutOutputDTO {
	workouts := use_case.WorkoutRepository.Find()

	var listWorkoutsOutput []*ListWorkoutOutputDTO

	if workouts == nil {
		listWorkoutsOutput = []*ListWorkoutOutputDTO{}

		return listWorkoutsOutput
	}

	for _, workout := range workouts {
		listWorkoutsOutput = append(listWorkoutsOutput, &ListWorkoutOutputDTO{
			ID:          workout.ID,
			Name:        workout.Name,
			Description: workout.Description,
			Exercises:   workout.Exercises,
			CreatedAt:   workout.CreatedAt,
		})
	}

	return listWorkoutsOutput
}
