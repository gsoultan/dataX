package mysql

import (
	"context"

	"github.com/gsoultan/dataX"
	"github.com/jmoiron/sqlx"
)

type transactionManager struct {
	db dataX.Database
	tx *sqlx.Tx
}

func (s *transactionManager) GetTransaction() interface{} {
	return s.tx
}

func (s *transactionManager) Begin(ctx context.Context) error {
	var err error

	s.tx, err = s.db.GetConnection().(*sqlx.DB).BeginTxx(ctx, nil)
	return err
}

func (s *transactionManager) Commit() error {
	return s.tx.Commit()
}

func (s *transactionManager) Rollback() error {
	return s.tx.Rollback()
}

func NewTransactionManager(db dataX.Database) dataX.TransactionManager {
	u := &transactionManager{}
	u.db = db
	return u
}
