package usecase

import (
	"context"

	"github.com/gofrs/uuid/v5"

	"matcha/internal/core"
	"matcha/internal/pkg/password"
)

type (
	CreateUserUsecase struct {
		r core.Repository
	}
	GetUserUsecase struct {
		r core.Repository
	}
)

func NewCreateUserUsecase(r core.Repository) *CreateUserUsecase {
	return &CreateUserUsecase{
		r: r,
	}
}

func (uc *CreateUserUsecase) Execute(ctx context.Context, args *core.CreateUserParams) (*core.User, error) {
	hashedPassword, err := password.Hash(args.Password)
	if err != nil {
		return nil, err
	}
	u := core.NewUser(args.Username, args.Email, hashedPassword)
	return uc.r.CreateUser(ctx, u)
}

func NewGetUserUsecase(r core.Repository) *GetUserUsecase {
	return &GetUserUsecase{
		r: r,
	}
}

func (uc *GetUserUsecase) Execute(ctx context.Context, id uuid.UUID) (*core.User, error) {
	return uc.r.GetUser(ctx, id)
}
