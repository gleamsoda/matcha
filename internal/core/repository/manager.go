package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/samber/do"

	"matcha/internal/core"
)

type (
	Manager struct {
		r TxRunner
		e Executor
	}
	TxRunner interface {
		BeginTx(context.Context, *sql.TxOptions) (*sql.Tx, error)
	}
	Executor interface {
		ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
		PrepareContext(context.Context, string) (*sql.Stmt, error)
		QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
		QueryRowContext(context.Context, string, ...interface{}) *sql.Row
	}
)

func NewManager(i *do.Injector) (core.RepositoryManager, error) {
	db := do.MustInvoke[*sql.DB](i)
	return &Manager{
		r: db,
		e: db,
	}, nil
}

func (m *Manager) User() core.UserRepository {
	return NewUser(m.e)
}

func (m *Manager) Issue() core.IssueRepository {
	return NewIssue(m.e)
}

func (m *Manager) Transaction(ctx context.Context, fn core.TransactionFunc) error {
	tx, err := m.r.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	txm := &Manager{r: m.r, e: tx}
	if err := fn(ctx, txm); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
