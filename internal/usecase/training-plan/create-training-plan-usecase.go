package usecase

import "github.com/mauFade/fit-force/internal/infra/repository"

type CreateTrainingPlanUseCase struct {
	TrainingRepository repository.TrainingPlanDatabaseRepository
}

func NewCreateTrainingPlanUseCase(repo repository.TrainingPlanDatabaseRepository) *CreateTrainingPlanUseCase {
	return &CreateTrainingPlanUseCase{
		TrainingRepository: repo,
	}
}
