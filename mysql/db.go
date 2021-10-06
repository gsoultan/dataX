package mysql

import (
	"context"
	"fmt"
	"sync"

	"github.com/jmoiron/sqlx"
)

var once sync.Once
var db *sqlx.DB

type database struct {
	driverName   string
	databaseName string
	uri          string
	ctx          context.Context
}

func (d *database) Ping() error {
	return db.Ping()
}

func (d *database) GetConnection() interface{} {
	if db != nil {
		return db
	}

	var err error
	var conn *sqlx.DB

	if conn, err = sqlx.ConnectContext(d.ctx, d.driverName, d.uri); err != nil {
		fmt.Println("error", err)
		return nil
	}

	once.Do(func() {
		db = conn
	})

	return db
}

func (d *database) GetDatabaseName() string {
	return d.databaseName
}

func (d *database) GetUri() string {
	return d.uri
}
