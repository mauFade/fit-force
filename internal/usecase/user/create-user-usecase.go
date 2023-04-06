package usecase

import (
	"github.com/mauFade/fit-force/internal/infra/repository"
	"github.com/mauFade/fit-force/internal/model"
)

type CreateUserInputDTO struct {
	Name     string
	Email    string
	Password string
	Age      int
	Gender   string
	Height   float64
	Weight   string
}

type CreateuserUseCase struct {
	UserRepository repository.UserDatabaseRepository
}

func NewCreateUserUseCase(repo repository.UserDatabaseRepository) *CreateuserUseCase {
	return &CreateuserUseCase{
		UserRepository: repo,
	}
}

func (use_case *CreateuserUseCase) Execute(data CreateUserInputDTO) (*model.User, error) {
	user := model.NewUser(
		data.Name,
		data.Email,
		data.Password,
		data.Age,
		data.Gender,
		data.Height,
		data.Weight,
	)

	err := user.Validate()

	if err != nil {
		return nil, err
	}

	use_case.UserRepository.Create(user)

	return user, nil
}
