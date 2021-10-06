package mongodb

import (
	"context"
	"fmt"

	"github.com/gsoultan/dataX"
	"github.com/gsoultan/dataX/builder"
)

type mongodbBuilder struct {
	config  *dataX.Config
	context context.Context
}

// WithContext returns a copy with the given context.Context context
func (b *mongodbBuilder) WithContext(context context.Context) builder.Database {
	b.context = context
	return b
}

func (b *mongodbBuilder) Build() dataX.Database {
	return &database{
		ctx:          b.context,
		databaseName: b.config.Database,
		uri:          b.createUri(),
	}
}

// WithConfig returns a copy with the given dataX.Config config
func (b *mongodbBuilder) WithConfig(config *dataX.Config) builder.Database {
	b.config = config
	return b
}

func NewBuilder() builder.Database {
	return &mongodbBuilder{}
}

func (b *mongodbBuilder) createUri() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", b.config.UserName, b.config.Password, b.config.Host, b.config.Port, b.config.Database)
}
