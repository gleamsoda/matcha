package repository

import (
	"context"

	"github.com/gofrs/uuid/v5"
	"github.com/stephenafamo/bob"

	"matcha/internal/core"
	sqlcgen "matcha/internal/core/repository/sqlc/gen"
)

type User struct {
	sqlc *sqlcgen.Queries
	bob  bob.Executor
}

func NewUser(db Executor) *User {
	return &User{
		sqlc: sqlcgen.New(db),
		bob:  bob.New(db),
	}
}

var _ core.UserRepository = (*User)(nil)

func (r *User) Create(ctx context.Context, u *core.User) (*core.User, error) {
	row, err := r.sqlc.CreateUser(ctx, &sqlcgen.CreateUserParams{
		Username:       u.Username,
		Email:          u.Email,
		HashedPassword: u.HashedPassword,
	})
	if err != nil {
		return nil, err
	}

	return &core.User{
		ID:             row.ID,
		Username:       row.Username,
		Email:          row.Email,
		HashedPassword: row.HashedPassword,
	}, err
}

func (r *User) Get(ctx context.Context, id uuid.UUID) (*core.User, error) {
	row, err := r.sqlc.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return &core.User{
		ID:             row.ID,
		Username:       row.Username,
		HashedPassword: row.HashedPassword,
		Email:          row.Email,
	}, nil
}
