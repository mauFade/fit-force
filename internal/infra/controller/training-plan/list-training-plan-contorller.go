package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mauFade/fit-force/internal/infra/repository"
	"github.com/mauFade/fit-force/internal/loaders"
	usecase "github.com/mauFade/fit-force/internal/usecase/training-plan"
)

func ListTrainingPlanController(context *gin.Context) {
	repo := repository.NewTrainigPlanRepository(loaders.DB)
	list_training_plan_usecase := usecase.NewListTrainingPlanUseCase(*repo)

	output := list_training_plan_usecase.Execute()

	context.JSON(200, output)
}
