package repository

import "gorm.io/gorm"

type WorkoutDatabaseRepository struct {
	DB *gorm.DB
}

func NewWorkoutRepository(database *gorm.DB) *WorkoutDatabaseRepository {
	return &WorkoutDatabaseRepository{
		DB: database,
	}
}
