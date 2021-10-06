package mongodb

import (
	"context"
	"fmt"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var once sync.Once
var client *mongo.Client

type database struct {
	databaseName string
	uri          string
	ctx          context.Context
}

func (d *database) Ping() error {
	return client.Ping(d.ctx, nil)
}

func (d *database) GetConnection() interface{} {
	if client != nil {
		return client
	}

	clientOptions := options.Client().ApplyURI(d.uri)

	var err error
	var cl *mongo.Client
	if cl, err = mongo.NewClient(clientOptions); err != nil {
		fmt.Println("mongo", "initiation", "err", err)
		return nil
	}
	if err = client.Connect(d.ctx); err != nil {
		fmt.Println("mongo", "connecting", "err", err)
		return nil
	}

	once.Do(func() {
		client = cl
	})
	return client
}

func (d *database) GetDatabaseName() string {
	return d.databaseName
}

func (d *database) GetUri() string {
	return d.uri
}
