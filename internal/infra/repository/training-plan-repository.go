package repository

import "gorm.io/gorm"

type TrainingPlanDatabaseRepository struct {
	DB *gorm.DB
}

func NewTrainigPlanRepository(database *gorm.DB) *TrainingPlanDatabaseRepository {
	return &TrainingPlanDatabaseRepository{
		DB: database,
	}
}
