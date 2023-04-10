package repository

import (
	"github.com/mauFade/fit-force/internal/model"
	"gorm.io/gorm"
)

type WorkoutDatabaseRepository struct {
	DB *gorm.DB
}

func NewWorkoutRepository(database *gorm.DB) *WorkoutDatabaseRepository {
	return &WorkoutDatabaseRepository{
		DB: database,
	}
}

func (repo *WorkoutDatabaseRepository) Create(workout *model.Workout) {
	repo.DB.Create(workout)
}

func (repo *WorkoutDatabaseRepository) Find() []*model.Workout {
	var workouts []*model.Workout

	repo.DB.Raw("SELECT * FROM workouts").Scan(&workouts)

	return workouts
}
