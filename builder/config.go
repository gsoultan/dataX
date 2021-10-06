package builder

import "github.com/gsoultan/dataX"

type Config interface {
	Build() *dataX.Config
	WithProvider(provider string) Config
	WithHost(host string) Config
	WithPort(port int) Config
	WithUserName(userName string) Config
	WithPassword(password string) Config
	WithDatabase(database string) Config
}

type config struct {
	provider string
	host     string
	port     int
	userName string
	password string
	database string
	build    *dataX.Config
}

// WithPort returns a copy with the given int port
func (c *config) WithPort(port int) Config {
	c.port = port
	return c
}

// GetBuild returns the *dataX.Config build
func (c *config) Build() *dataX.Config {
	return c.build
}

// WithProvider returns a copy with the given string provider
func (c *config) WithProvider(provider string) Config {
	c.provider = provider
	return c
}

// WithHost returns a copy with the given string host
func (c *config) WithHost(host string) Config {
	c.host = host
	return c
}

// WithUserName returns a copy with the given string userName
func (c *config) WithUserName(userName string) Config {
	c.userName = userName
	return c
}

// WithPassword returns a copy with the given string password
func (c *config) WithPassword(password string) Config {
	c.password = password
	return c
}

// WithDatabase returns a copy with the given string database
func (c *config) WithDatabase(database string) Config {
	c.database = database
	return c
}

func NewConfig() Config {
	return &config{}
}
