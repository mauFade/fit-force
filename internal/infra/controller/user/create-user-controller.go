package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mauFade/fit-force/internal/infra/repository"
	"github.com/mauFade/fit-force/internal/loaders"
	usecase "github.com/mauFade/fit-force/internal/usecase/user"
)

type CreateUserInputDTO struct {
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Age      int     `json:"age"`
	Gender   string  `json:"gender"`
	Height   float64 `json:"height"`
	Weight   string  `json:"weight"`
}

func CreateUserController(context *gin.Context) {
	var body CreateUserInputDTO

	context.Bind(&body)

	repo := repository.NewUserRepository(loaders.DB)
	create_user_usecase := usecase.NewCreateUserUseCase(*repo)

	output, err := create_user_usecase.Execute(usecase.CreateUserInputDTO{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
		Age:      body.Age,
		Gender:   body.Gender,
		Height:   body.Height,
		Weight:   body.Weight,
	})

	if err != nil {
		context.JSON(403, gin.H{
			"error": err.Error(),
		})

		return
	}

	context.JSON(201, output)
}
