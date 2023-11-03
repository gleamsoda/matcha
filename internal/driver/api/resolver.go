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
	issues     []*core.Issue
	getUser    core.GetUserUsecase
	createUser core.CreateUserUsecase
}

func NewResolver(i *do.Injector) (*Resolver, error) {
	r := do.MustInvoke[core.Repository](i)
	return &Resolver{
		getUser:    usecase.NewGetUserUsecase(r),
		createUser: usecase.NewCreateUserUsecase(r),
	}, nil
}
