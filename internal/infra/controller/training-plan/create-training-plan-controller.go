package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mauFade/fit-force/internal/infra/repository"
	"github.com/mauFade/fit-force/internal/loaders"
	usecase "github.com/mauFade/fit-force/internal/usecase/training-plan"
)

type CreateTrainingPLanInputDTO struct {
	UserID      string `json:"user_id"`
	WorkoutID   string `json:"workout_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Objective   string `json:"objective"`
	Duration    string `json:"duration"`
}

func CreateTrainingPlanController(context *gin.Context) {
	var body CreateTrainingPLanInputDTO

	context.Bind(&body)

	repo := repository.NewTrainigPlanRepository(loaders.DB)
	create_trainingplan_usecase := usecase.NewCreateTrainingPlanUseCase(*repo)

	output, err := create_trainingplan_usecase.Execute(usecase.CreateTrainingPLanInputDTO{
		UserID:      body.UserID,
		WorkoutID:   body.WorkoutID,
		Name:        body.Name,
		Description: body.Description,
		Objective:   body.Objective,
		Duration:    body.Duration,
	})

	if err != nil {
		context.JSON(403, gin.H{
			"error": err.Error(),
		})

		return
	}

	context.JSON(201, output)
}
