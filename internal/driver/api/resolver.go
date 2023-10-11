package api

import (
	"github.com/samber/do"

	"matcha/internal/core"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	users  []*core.User
	issues []*core.Issue
}

func NewResolver(i *do.Injector) (*Resolver, error) {
	return &Resolver{}, nil
}
