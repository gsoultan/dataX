package mongodb

import (
	"context"
	"errors"

	"github.com/gsoultan/dataX"
	"go.mongodb.org/mongo-driver/mongo"
)

type transactionManager struct {
	db      dataX.Database
	session mongo.Session
	ctx     context.Context
}

func (t *transactionManager) Begin(ctx context.Context) error {
	var err error

	if t.db == nil {
		return errors.New("database is not instantiated")
	}

	if t.session, err = t.db.GetConnection().(*mongo.Client).StartSession(); err != nil {
		return err
	}
	if err = t.session.StartTransaction(); err != nil {
		return err
	}

	t.ctx = ctx
	return nil
}

func (t *transactionManager) Commit() error {
	if err := t.session.CommitTransaction(t.ctx); err != nil {
		t.session.EndSession(t.ctx)
		return err
	}
	t.session.EndSession(t.ctx)
	return nil
}

func (t *transactionManager) Rollback() error {
	if err := t.session.AbortTransaction(t.ctx); err != nil {
		t.session.EndSession(t.ctx)
		return err
	}
	t.session.EndSession(t.ctx)
	return nil
}

func (t *transactionManager) GetTransaction() interface{} {
	return t.session
}

func NewTransactionManager(db dataX.Database) dataX.TransactionManager {
	return &transactionManager{db: db}
}
