package repository

import (
	"errors"

	"github.com/mauFade/fit-force/internal/model"
	"gorm.io/gorm"
)

type UserDatabaseRepository struct {
	DB *gorm.DB
}

func NewUserRepository(database *gorm.DB) *UserDatabaseRepository {
	return &UserDatabaseRepository{
		DB: database,
	}
}

func (repo *UserDatabaseRepository) Create(user *model.User) {
	repo.DB.Create(user)
}

func (repo *UserDatabaseRepository) FindById(user_id string) (*model.User, error) {
	var user *model.User

	result := repo.DB.Raw("SELECT * FROM users WHERE id = ?", user_id).Scan(&user)

	if result == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (repo *UserDatabaseRepository) FindByEmail(user_email string) (*model.User, error) {
	var user model.User

	if err := repo.DB.Where("email = ?", user_email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}
