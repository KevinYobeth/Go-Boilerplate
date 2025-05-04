package user

import (
	"github.com/google/uuid"
)

type UserDto struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}

func NewUserDto(firstName, lastName, email, password string) *UserDto {
	return &UserDto{
		ID:        uuid.New(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
	}
}
