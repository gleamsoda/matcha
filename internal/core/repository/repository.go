package repository

import (
	"context"
	"database/sql"

	"github.com/samber/do"

	"matcha/internal/core"
)

type (
	Repository struct {
		exec Executor
		txn  core.Transaction
	}
	DB interface {
		BeginTx(context.Context, *sql.TxOptions) (*sql.Tx, error)
		Executor
	}
	Executor interface {
		ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
		PrepareContext(context.Context, string) (*sql.Stmt, error)
		QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
		QueryRowContext(context.Context, string, ...interface{}) *sql.Row
	}
)

func NewRepository(i *do.Injector) (core.Repository, error) {
	db := do.MustInvoke[*sql.DB](i)
	return &Repository{
		txn:  NewTransaction(db),
		exec: db,
	}, nil
}

func (m *Repository) User() core.UserRepository {
	return NewUser(m.exec)
}

func (m *Repository) Issue() core.IssueRepository {
	return NewIssue(m.exec)
}

func (m *Repository) Transaction() core.Transaction {
	return m.txn
}
