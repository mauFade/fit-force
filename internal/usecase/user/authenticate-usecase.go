package usecase

import (
	"errors"

	"github.com/mauFade/fit-force/internal/infra/auth"
	"github.com/mauFade/fit-force/internal/infra/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticateInputDTO struct {
	Email    string
	Password string
}

type AuthenticateOutputDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type AuthenticateUseCase struct {
	UserRepository repository.UserDatabaseRepository
}

func NewAuthenticateUseCase(repo repository.UserDatabaseRepository) *AuthenticateUseCase {
	return &AuthenticateUseCase{
		UserRepository: repo,
	}
}

func (use_case *AuthenticateUseCase) Execute(data AuthenticateInputDTO) (*AuthenticateOutputDTO, error) {
	if data.Email == "" || data.Password == "" {
		return nil, errors.New("email or password not provided")
	}

	user := use_case.UserRepository.FindByEmail(data.Email)

	if user == nil {
		return nil, errors.New("user not found with this email")
	}

	passwordErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))

	if passwordErr != nil {
		return nil, errors.New("invalid password")
	}

	// Gera um token JWT
	user_token, err := auth.GenerateToken(user.ID)

	if err != nil {
		return nil, err
	}

	return &AuthenticateOutputDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: user_token,
	}, nil
}
