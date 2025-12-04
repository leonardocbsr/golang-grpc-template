package config

import (
	"fmt"

	"github.com/leonardocbsr/golang-grpc-template/internal/errors"
	"github.com/spf13/viper"
)

var _ Config = &config{}

type config struct {
	Database DatabaseConfig `mapstructure:"database"`
	Logger   LoggerConfig   `mapstructure:"logger"`
	Server   ServerConfig   `mapstructure:"server"`
}

// New loads the configuration from file and environment variables.
func New() (Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var instance config
	if err := viper.Unmarshal(&instance); err != nil {
		return nil, fmt.Errorf("failed to decode config: %w", err)
	}

	return &instance, nil
}

func (c *config) GetDatabaseConfig() DatabaseConfig {
	return c.Database
}

func (c *config) GetLoggerConfig() LoggerConfig {
	return c.Logger
}

func (c *config) GetServerConfig() ServerConfig {
	return c.Server
}

func (c *config) GetClientsConfig() []ClientsConfig {
	return c.Server.Clients
}

func (c *config) GetServerURL() string {
	return fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port)
}

func (c *config) GetClientConfig(name string) (*ClientsConfig, error) {
	for _, client := range c.Server.Clients {
		if client.Name == name {
			return &client, nil
		}
	}

	return nil, errors.ErrInvalidClientName
}
