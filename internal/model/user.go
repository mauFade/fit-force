package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mauFade/fit-force/internal/formatter"
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
	feet, inches := formatter.CMToFeet(height)

	heightStr := fmt.Sprintf("%d'%d\"", feet, inches)

	return &User{
		ID:       uuid.NewString(),
		Name:     name,
		Email:    email,
		Password: password,
		Age:      age,
		Gender:   gender,
		Height:   heightStr,
		Weight:   weight,
	}
}
