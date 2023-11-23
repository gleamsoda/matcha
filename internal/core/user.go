package core

import (
	"context"

	"github.com/gofrs/uuid/v5"
)

type User struct {
	ID             uuid.UUID `json:"id"`
	Username       string    `json:"username"`
	HashedPassword string    `json:"-"`
	Email          string    `json:"email"`
}

func NewUser(username, email, hashedPassword string) *User {
	return &User{
		Username:       username,
		HashedPassword: hashedPassword,
		Email:          email,
	}
}

type (
	CreateUserUsecase interface {
		Execute(ctx context.Context, args *CreateUserParams) (*User, error)
	}
	CreateUserParams struct {
		Username string
		Email    string
		Password string
	}
	GetUserUsecase interface {
		Execute(ctx context.Context, id uuid.UUID) (*User, error)
	}
)
