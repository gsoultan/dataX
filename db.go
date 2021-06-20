package dataX

type Database interface {
	GetDatabaseName() string
	GetUri() string
	GetConnection() (interface{}, error)
	Ping() error
}
