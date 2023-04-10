package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mauFade/fit-force/internal/infra/repository"
	"github.com/mauFade/fit-force/internal/loaders"
	usecase "github.com/mauFade/fit-force/internal/usecase/user"
)

type AuthenticateInputDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AuthenticateController(context *gin.Context) {
	var body AuthenticateInputDTO

	context.BindJSON(&body)

	repo := repository.NewUserRepository(loaders.DB)
	use_case := usecase.NewAuthenticateUseCase(*repo)

	output, err := use_case.Execute(usecase.AuthenticateInputDTO{
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		context.JSON(401, gin.H{
			"error": err.Error(),
		})

		return
	}

	context.JSON(200, output)
}
