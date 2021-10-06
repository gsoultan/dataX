package builder

import (
	"context"

	"github.com/gsoultan/dataX"
)

type Database interface {
	Build() dataX.Database
	WithConfig(config dataX.Config) Database
	WithContext(ctx context.Context) Database
}
