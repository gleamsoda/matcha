package repository

import (
	"context"
	"database/sql"

	"github.com/gofrs/uuid/v5"
	"github.com/samber/do"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"

	"matcha/internal/core"
	"matcha/internal/core/repository/sqlc"
)

type Repository struct {
	db   *sql.DB
	bun  *bun.DB
	sqlc *sqlc.Queries
}

func NewRepository(i *do.Injector) (core.Repository, error) {
	db := do.MustInvoke[*sql.DB](i)
	return &Repository{
		db:   db,
		bun:  bun.NewDB(db, pgdialect.New()),
		sqlc: sqlc.New(db),
	}, nil
}

func (r *Repository) CreateUser(ctx context.Context, u *core.User) (*core.User, error) {
	row, err := r.sqlc.CreateUser(ctx, &sqlc.CreateUserParams{
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

func (r *Repository) GetUser(ctx context.Context, id uuid.UUID) (*core.User, error) {
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
