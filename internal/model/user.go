package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mauFade/fit-force/internal/utils"
)

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	Age       int
	Gender    string
	Height    string
	Weight    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(
	name string,
	email string,
	password string,
	age int,
	gender string,
	height float64, // Recebe a altura em cm e usa uma função formatadora para salvar em pés
	weight string,
) *User {
	feet, inches := utils.CMToFeet(height)

	heightStr := fmt.Sprintf("%d'%d", feet, inches)

	return &User{
		ID:        uuid.NewString(),
		Name:      name,
		Email:     email,
		Password:  password,
		Age:       age,
		Gender:    gender,
		Height:    heightStr,
		Weight:    weight,
		CreatedAt: time.Now(),
	}
}

func (user *User) Validate() error {
	if user.ID == "" {
		return errors.New("user ID is requried")
	}

	if user.Name == "" {
		return errors.New("user name is required")
	}

	if user.Email == "" || !utils.ValidateMail(user.Email) {
		return errors.New("invalid or null email, please provide a valid one")
	}

	if user.Password == "" {
		return errors.New("user password is required")
	}

	if user.Age <= 0 {
		return errors.New("user age is required")
	}

	if user.Gender == "" || !utils.Contains([]string{"male", "female", "non-specified"}, user.Gender) {
		return errors.New("invalid user gender")
	}

	if user.Height == "" {
		return errors.New("user height is required")
	}

	if user.Weight == "" {
		return errors.New("user weight is required")
	}

	return nil
}

type UserRepository interface {
	Create(user *User) error
	FindById(user_id string) (*User, error)
	FindByEmail(user_email string) *User
}
