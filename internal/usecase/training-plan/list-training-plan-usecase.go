package usecase

import (
	"time"

	"github.com/mauFade/fit-force/internal/infra/repository"
)

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

type ListTrainingPlanUseCase struct {
	TrainigPlanRepository repository.TrainingPlanDatabaseRepository
}

func NewListTrainingPlanUseCase(repo repository.TrainingPlanDatabaseRepository) *ListTrainingPlanUseCase {
	return &ListTrainingPlanUseCase{
		TrainigPlanRepository: repo,
	}
}

func (use_case *ListTrainingPlanUseCase) Execute() []*ListTrainingPLanOutputDTO {
	training_plans := use_case.TrainigPlanRepository.Find()

	var training_plans_output []*ListTrainingPLanOutputDTO

	if training_plans == nil {
		training_plans_output = []*ListTrainingPLanOutputDTO{}

		return training_plans_output
	}

	for _, training_plan := range training_plans {
		training_plans_output = append(training_plans_output, &ListTrainingPLanOutputDTO{
			ID:          training_plan.ID,
			UserID:      training_plan.UserID,
			WorkoutID:   training_plan.WorkoutID,
			Name:        training_plan.Name,
			Description: training_plan.Description,
			Objective:   training_plan.Objective,
			Duration:    training_plan.Duration,
			CreatedAt:   training_plan.CreatedAt,
		})
	}

	return training_plans_output
}
