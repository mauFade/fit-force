package usecase

import (
	"encoding/json"
	"time"

	"github.com/mauFade/fit-force/internal/infra/repository"
	"github.com/mauFade/fit-force/internal/model"
)

type CreateWorkoutInputDTO struct {
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
	jsonBytes, err := json.Marshal(data.Exercises)

	if err != nil {
		return nil, err
	}

	exercises := string(jsonBytes)

	workout := model.NewWorkout(data.Name, data.Description, exercises)

	err = workout.Validate()

	if err != nil {
		return nil, err
	}

	use_case.WorkoutRepository.Create(workout)

	var exercisesArr []string

	err = json.Unmarshal([]byte(exercises), &exercisesArr)

	if err != nil {
		return nil, err
	}

	return &CreateWorkoutOutputDTO{
		ID:          workout.ID,
		Name:        workout.Name,
		Description: workout.Description,
		Exercises:   exercisesArr,
		CreatedAt:   workout.CreatedAt,
	}, nil
}
