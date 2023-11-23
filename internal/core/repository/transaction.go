package repository

import (
	"context"
	"fmt"

	"matcha/internal/core"
)

type (
	Transaction            struct{ db DB }
	TransactionUnsupported struct{}
)

func NewTransaction(db DB) *Transaction {
	return &Transaction{db: db}
}

var _ core.Transaction = (*Transaction)(nil)

func (t *Transaction) Run(ctx context.Context, fn core.TransactionFunc) error {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	txm := &Repository{exec: tx, txn: &TransactionUnsupported{}}
	if err := fn(ctx, txm); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

func (t *TransactionUnsupported) Run(ctx context.Context, fn core.TransactionFunc) error {
	panic("nested transactions are not supported")
}
