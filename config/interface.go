package config

import "fmt"

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"name"`
}

type LoggerConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

type ClientsConfig struct {
	Name string `mapstructure:"name"`
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func (c ClientsConfig) GetURL() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

type ServerConfig struct {
	Port    int             `mapstructure:"port"`
	Host    string          `mapstructure:"host"`
	Clients []ClientsConfig `mapstructure:"clients"`
}

type IConfig interface {
	// GetDatabaseConfig returns the database configuration
	GetDatabaseConfig() DatabaseConfig
	// GetLoggerConfig returns the logger configuration
	GetLoggerConfig() LoggerConfig
	// GetServerConfig returns the server configuration
	GetServerConfig() ServerConfig
	// GetClientsConfig returns the clients configuration
	GetClientsConfig() []ClientsConfig

	// GetServerURL returns the server URL
	GetServerURL() string

	// GetClientConfig returns the client configuration
	GetClientConfig(name string) (*ClientsConfig, error)
}
