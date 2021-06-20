package sqlx

import (
	"fmt"
	"github.com/gsoultan/dataX"
	"github.com/jmoiron/sqlx"
)

type database struct {
	db           *sqlx.DB
	driverName   string
	databaseName string
	uri          string
}

func (d *database) Ping() error {
	return d.db.Ping()
}

func (d *database) GetConnection() (interface{}, error) {
	if d.db != nil {
		return d.db, nil
	}

	var err error
	if d.db, err = sqlx.Connect(d.driverName, d.GetUri()); err != nil {
		return nil, err
	}
	return d.db, nil
}

func (d *database) GetDatabaseName() string {
	return d.databaseName
}

func (d *database) GetUri() string {
	return d.uri
}

func connect(driverName string, uri string) (*sqlx.DB, error) {
	return sqlx.Connect(driverName, uri)
}

func New(driverName string, cfg dataX.Config) (dataX.Database, error) {
	var err error

	d := &database{}
	d.databaseName = cfg.Database

	d.uri = createUri(driverName, cfg)
	if d.db, err = connect(driverName, d.GetUri()); err != nil {
		return nil, err
	}

	return d, nil
}

func createUri(driverName string, cfg dataX.Config) string {
	if driverName == SQLX_DRIVER_MYSQL {
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.UserName, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	}
	return ""
}
