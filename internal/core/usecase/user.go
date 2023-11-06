package usecase

import (
	"context"

	"github.com/gofrs/uuid/v5"

	"matcha/internal/core"
	"matcha/internal/pkg/password"
)

type (
	CreateUser struct {
		r core.RepositoryManager
	}
	GetUser struct {
		r core.RepositoryManager
	}
)

func NewCreateUser(r core.RepositoryManager) *CreateUser {
	return &CreateUser{
		r: r,
	}
}

var _ core.CreateUserUsecase = (*CreateUser)(nil)

func (uc *CreateUser) Execute(ctx context.Context, args *core.CreateUserParams) (*core.User, error) {
	hashedPassword, err := password.Hash(args.Password)
	if err != nil {
		return nil, err
	}
	u := core.NewUser(args.Username, args.Email, hashedPassword)
	return uc.r.User().Create(ctx, u)
}

var _ core.GetUserUsecase = (*GetUser)(nil)

func NewGetUser(r core.RepositoryManager) *GetUser {
	return &GetUser{
		r: r,
	}
}

func (uc *GetUser) Execute(ctx context.Context, id uuid.UUID) (*core.User, error) {
	return uc.r.User().Get(ctx, id)
}
