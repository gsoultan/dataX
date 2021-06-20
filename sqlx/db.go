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
	return sqlx.Connect(d.driverName, d.GetUri())
}

func (d *database) GetDatabaseName() string {
	return d.databaseName
}

func (d *database) GetUri() string {
	return d.uri
}

func New(driverName string, cfg dataX.Config) dataX.Database {
	d := &database{}
	d.databaseName = cfg.Database

	d.uri = createUri(driverName, cfg)
	return d
}

func createUri(driverName string, cfg dataX.Config) string {
	if driverName == SQLX_DRIVER_MYSQL {
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.UserName, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	}
	return ""
}
