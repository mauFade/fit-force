package loaders

import (
	"os"

	"github.com/mauFade/fit-force/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error

	connectionURI := os.Getenv("DATABASE_URL")

	DB, err = gorm.Open(postgres.Open(connectionURI), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// Migrations
	DB.AutoMigrate(
		&model.User{},
		&model.TrainingPlan{},
		&model.Workout{},
	)
}
