package dataX

import "context"

type TransactionManager interface {
	Begin(ctx context.Context) error
	Commit() error
	Rollback() error
	GetTransaction() interface{}
}
