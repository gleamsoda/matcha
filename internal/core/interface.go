package core

import (
	"context"

	"github.com/gofrs/uuid/v5"
)

type Repository interface {
	CreateUser(ctx context.Context, u *User) (*User, error)
	GetUser(ctx context.Context, id uuid.UUID) (*User, error)
}
