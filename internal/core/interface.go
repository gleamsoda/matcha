package core

import "context"

type (
	RepositoryManager interface {
		User() UserRepository
		Issue() IssueRepository
	}
	TransactionFunc func(context.Context, RepositoryManager) error
)
