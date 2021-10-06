package factory

import (
	"context"

	"github.com/gsoultan/builder"
	"github.com/gsoultan/dataX"
	"github.com/gsoultan/dataX/flags"
	"github.com/gsoultan/dataX/mongodb"
	"github.com/gsoultan/dataX/mysql"
)

func Create(ctx context.Context, b builder.Config) dataX.Database {
	config := b.Build()

	if config.Provider == flags.DRIVER_MYSQL {
		mBuilder := mysql.NewBuilder()
		mBuilder.WithConfig(config).WithContext(ctx)
		return mBuilder.Build()
	}

	if config.Provider == flags.DRIVER_MONGODB {
		mdBuilder := mongodb.NewBuilder()
		mdBuilder.WithConfig(config).WithContext(ctx)
		return mdBuilder.Build()
	}
	return nil
}
