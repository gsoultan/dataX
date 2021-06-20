package sqlx

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
	var db interface{}

	if db, err = s.db.GetConnection(); err != nil {
		return err
	}

	s.tx, err = db.(*sqlx.DB).BeginTxx(ctx, nil)
	return err
}

func (s *transactionManager) Commit() error {
	return s.Commit()
}

func (s *transactionManager) Rollback() error {
	return s.Rollback()
}

func NewTransactionManager(db dataX.Database) dataX.TransactionManager {
	u := &transactionManager{}
	u.db = db
	return u
}
