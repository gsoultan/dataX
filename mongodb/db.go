package mongodb

import (
	"context"
	"fmt"
	"github.com/gsoultan/dataX"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type database struct {
	databaseName string
	uri          string
	client       *mongo.Client
}

func (d *database) GetConnection() *mongo.Client {
	return d.client
}

func (d *database) GetDatabaseName() string {
	return d.databaseName
}

func (d *database) GetUri() string {
	return d.uri
}

func New(ctx context.Context, cfg dataX.Config) (dataX.Database, error) {
	u := &database{}
	u.databaseName = cfg.Database
	u.uri = fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", cfg.UserName, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	clientOptions := options.Client().ApplyURI(u.GetUri())

	var err error
	if u.client, err = mongo.NewClient(clientOptions); err != nil {
		return nil, err
	}

	u.client.Connect(ctx)
	if err = u.client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return u, nil
}
