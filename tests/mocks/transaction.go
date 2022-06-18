package mocks

import (
	"context"
)

type Transaction struct {
	MockRollbackFn func(ctx context.Context) error
	MockExecFn     func(ctx context.Context, sql string, arguments ...interface{}) (int64, error)
	MockCommitFn   func(ctx context.Context) error
}

func (t *Transaction) Rollback(ctx context.Context) error {
	return t.MockRollbackFn(ctx)
}

func (t *Transaction) Exec(ctx context.Context, sql string, arguments ...interface{}) (int64, error) {
	return t.MockExecFn(ctx, sql, arguments)
}

func (t *Transaction) Commit(ctx context.Context) error {
	return t.MockCommitFn(ctx)
}

func NewTransaction() *Transaction {
	return &Transaction{
		MockCommitFn: func(ctx context.Context) error {
			return nil
		},
		MockExecFn: func(ctx context.Context, sql string, arguments ...interface{}) (int64, error) {
			return 1, nil
		},
		MockRollbackFn: func(ctx context.Context) error {
			return nil
		},
	}
}
