package core

import (
	"context"

	"github.com/gofrs/uuid/v5"
)

type (
	Repository interface {
		Transaction() Transaction
		User() UserRepository
		Issue() IssueRepository
	}
	Transaction interface {
		Run(ctx context.Context, fn TransactionFunc) error
	}
	TransactionFunc func(context.Context, Repository) error
	UserRepository  interface {
		Create(ctx context.Context, u *User) (*User, error)
		Get(ctx context.Context, id uuid.UUID) (*User, error)
	}
	IssueRepository interface {
		Create(ctx context.Context, u *Issue) (*Issue, error)
		List(ctx context.Context) ([]*Issue, error)
	}
)
