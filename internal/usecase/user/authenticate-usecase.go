package usecase

import "github.com/mauFade/fit-force/internal/infra/repository"

type AuthenticateUseCase struct {
	UserRepository repository.UserDatabaseRepository
}
