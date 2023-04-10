package usecase

import (
	"time"

	"github.com/mauFade/fit-force/internal/infra/repository"
	"github.com/mauFade/fit-force/internal/model"
)

type CreateWorkoutInputDTO struct {
	ID          string
	Name        string
	Description string
	Exercises   []string
}

type CreateWorkoutOutputDTO struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Exercises   []string  `json:"exercises"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateWorkoutUseCase struct {
	WorkoutRepository repository.WorkoutDatabaseRepository
}

func NewCreateWorkoutUseCase(repo repository.WorkoutDatabaseRepository) *CreateWorkoutUseCase {
	return &CreateWorkoutUseCase{
		WorkoutRepository: repo,
	}
}

func (use_case *CreateWorkoutUseCase) Execute(data CreateWorkoutInputDTO) (*CreateWorkoutOutputDTO, error) {
	workout := model.NewWorkout(data.Name, data.Description, data.Exercises)

	err := workout.Validate()

	if err != nil {
		return nil, err
	}

	use_case.WorkoutRepository.Create(workout)

	return &CreateWorkoutOutputDTO{
		ID:          workout.ID,
		Name:        workout.Name,
		Description: workout.Description,
		Exercises:   workout.Exercises,
		CreatedAt:   workout.CreatedAt,
	}, nil
}
