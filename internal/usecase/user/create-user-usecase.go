package usecase

import (
	"errors"
	"time"

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

type CreateUserOutputDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	Gender    string    `json:"gender"`
	Height    string    `json:"height"`
	Weight    string    `json:"weight"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateuserUseCase struct {
	UserRepository repository.UserDatabaseRepository
}

func NewCreateUserUseCase(repo repository.UserDatabaseRepository) *CreateuserUseCase {
	return &CreateuserUseCase{
		UserRepository: repo,
	}
}

func (use_case *CreateuserUseCase) Execute(data CreateUserInputDTO) (*CreateUserOutputDTO, error) {
	email_already_exist := use_case.UserRepository.FindByEmail(data.Email)

	if email_already_exist != nil {
		return nil, errors.New("this email is already in use")
	}

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

	return &CreateUserOutputDTO{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Age:       user.Age,
		Gender:    user.Gender,
		Height:    user.Height,
		Weight:    user.Weight,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
