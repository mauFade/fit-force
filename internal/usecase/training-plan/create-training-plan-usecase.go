package usecase

import (
	"time"

	"github.com/mauFade/fit-force/internal/infra/repository"
	"github.com/mauFade/fit-force/internal/model"
)

type CreateTrainingPLanInputDTO struct {
	UserID      string
	WorkoutID   string
	Name        string
	Description string
	Objective   string
	Duration    string
}

type CreateTrainingPLanOutputDTO struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	WorkoutID   string    `json:"workout_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Objective   string    `json:"objective"`
	Duration    string    `json:"duration"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateTrainingPlanUseCase struct {
	TrainingRepository repository.TrainingPlanDatabaseRepository
}

func NewCreateTrainingPlanUseCase(repo repository.TrainingPlanDatabaseRepository) *CreateTrainingPlanUseCase {
	return &CreateTrainingPlanUseCase{
		TrainingRepository: repo,
	}
}

func (use_case *CreateTrainingPlanUseCase) Execute(data CreateTrainingPLanInputDTO) (*CreateTrainingPLanOutputDTO, error) {
	training_plan := model.NewTrainingPLan(
		data.Name,
		data.UserID,
		data.WorkoutID,
		data.Description,
		data.Objective,
		data.Duration,
	)

	err := training_plan.Validate()

	if err != nil {
		return nil, err
	}

	use_case.TrainingRepository.Create(training_plan)

	return &CreateTrainingPLanOutputDTO{
		ID:          training_plan.ID,
		UserID:      training_plan.UserID,
		WorkoutID:   training_plan.WorkoutID,
		Name:        training_plan.Name,
		Description: training_plan.Description,
		Objective:   training_plan.Objective,
		Duration:    training_plan.Duration,
		CreatedAt:   training_plan.CreatedAt,
	}, nil
}
