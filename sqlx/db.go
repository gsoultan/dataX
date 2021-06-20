package sqlx

import (
	"fmt"
	"github.com/gsoultan/dataX"
)

type database struct {
	databaseName string
	uri          string
}

func (d *database) GetDatabaseName() string {
	return d.databaseName
}

func (d *database) GetUri() string {
	return d.uri
}

func New(cfg dataX.Config) dataX.Database {
	d := &database{}
	d.databaseName = cfg.Database
	d.uri = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.UserName, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	return d
}
