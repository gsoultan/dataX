package mysql

import (
	"context"
	"fmt"

	"github.com/gsoultan/dataX"
	"github.com/gsoultan/dataX/builder"
)

type mysqlBuilder struct {
	config dataX.Config
	ctx    context.Context
}

// WithConfig returns a copy with the given dataX.Config config
func (b *mysqlBuilder) WithConfig(config dataX.Config) builder.Database {
	b.config = config
	return b
}

func (b *mysqlBuilder) Build() dataX.Database {
	return &database{
		driverName:   b.config.Provider,
		uri:          b.createUri(),
		databaseName: b.config.Database,
		ctx:          b.ctx,
	}
}

// WithContext returns a copy with the given context.Context ctx
func (b *mysqlBuilder) WithContext(ctx context.Context) builder.Database {
	b.ctx = ctx
	return b
}

func (b *mysqlBuilder) createUri() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", b.config.UserName, b.config.Password, b.config.Host, b.config.Port, b.config.Database)
}

func NewBuilder() builder.Database {
	return &mysqlBuilder{}
}
