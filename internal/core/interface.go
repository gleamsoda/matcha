package core

import "context"

type (
	RepositoryManager interface {
		User() UserRepository
	}
	TransactionFunc func(context.Context, RepositoryManager) error
)
