package sqlx

import (
	"context"
	"github.com/gsoultan/dataX"
	"github.com/jmoiron/sqlx"
)

type transactionManager struct {
	db *sqlx.DB
	tx *sqlx.Tx
}

func (s *transactionManager) GetTransaction() interface{} {
	return s.tx
}

func (s *transactionManager) Begin(ctx context.Context) error {
	var err error
	s.tx, err = s.db.BeginTxx(ctx, nil)
	return err
}

func (s *transactionManager) Commit() error {
	return s.Commit()
}

func (s *transactionManager) Rollback() error {
	return s.Rollback()
}

func NewTransactionManager(db *sqlx.DB) dataX.TransactionManager {
	u := &transactionManager{}
	u.db = db
	return u
}
