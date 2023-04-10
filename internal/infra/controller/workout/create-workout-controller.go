package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mauFade/fit-force/internal/infra/repository"
	"github.com/mauFade/fit-force/internal/loaders"
	usecase "github.com/mauFade/fit-force/internal/usecase/workout"
)

type CreateWorkoutInputDTO struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Exercises   []string `json:"exercises"`
}

func CreateWorkoutController(context *gin.Context) {
	var body *CreateWorkoutInputDTO

	context.BindJSON(&body)

	repo := repository.NewWorkoutRepository(loaders.DB)
	use_case := usecase.NewCreateWorkoutUseCase(*repo)

	output, err := use_case.Execute(usecase.CreateWorkoutInputDTO{
		Name:        body.Name,
		Description: body.Description,
		Exercises:   body.Exercises,
	})

	if err != nil {
		context.JSON(403, gin.H{
			"error": err.Error(),
		})

		return
	}

	context.JSON(201, output)
}
