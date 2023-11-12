package api

import (
	"github.com/samber/do"

	"matcha/internal/core"
	"matcha/internal/core/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	getUser     core.GetUserUsecase
	createUser  core.CreateUserUsecase
	createIssue core.CreateIssueUsecase
	listIssues  core.ListIssuesUsecase
}

func NewResolver(i *do.Injector) (*Resolver, error) {
	r := do.MustInvoke[core.RepositoryManager](i)
	return &Resolver{
		getUser:     usecase.NewGetUser(r),
		createUser:  usecase.NewCreateUser(r),
		createIssue: usecase.NewCreateIssue(r),
		listIssues:  usecase.NewListIssues(r),
	}, nil
}
