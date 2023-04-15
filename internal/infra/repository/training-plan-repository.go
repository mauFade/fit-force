package repository

import (
	"github.com/mauFade/fit-force/internal/model"
	"gorm.io/gorm"
)

type TrainingPlanDatabaseRepository struct {
	DB *gorm.DB
}

func NewTrainigPlanRepository(database *gorm.DB) *TrainingPlanDatabaseRepository {
	return &TrainingPlanDatabaseRepository{
		DB: database,
	}
}

func (repo *TrainingPlanDatabaseRepository) Create(training_plan *model.TrainingPlan) {
	repo.DB.Create(training_plan)
}

func (repo *TrainingPlanDatabaseRepository) Find() []*model.TrainingPlan {
	var trainings []*model.TrainingPlan

	repo.DB.Raw("SELECT * FROM training_plans").Scan(&trainings)

	return trainings
}
