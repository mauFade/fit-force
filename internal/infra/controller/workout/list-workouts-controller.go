package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mauFade/fit-force/internal/infra/repository"
	"github.com/mauFade/fit-force/internal/loaders"
	usecase "github.com/mauFade/fit-force/internal/usecase/workout"
)

func ListWorkoutsController(context *gin.Context) {
	repo := repository.NewWorkoutRepository(loaders.DB)
	list_workouts_usecase := usecase.NewListWorkoutsUseCase(*repo)

	output := list_workouts_usecase.Execute()

	context.JSON(200, output)
}
