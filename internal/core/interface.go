package core

import "context"

type (
	Repository interface {
		User() UserRepository
		Issue() IssueRepository
	}
	Transaction interface {
		Run(ctx context.Context, fn TransactionFunc) error
	}
	TransactionFunc func(context.Context, Repository) error
)
