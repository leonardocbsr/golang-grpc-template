package config

import (
	"fmt"

	"cbsr.io/golang-grpc-template/common/exceptions"
	"github.com/spf13/viper"
)

var _ IConfig = &config{}

type config struct {
	Database DatabaseConfig `mapstructure:"database"`
	Logger   LoggerConfig   `mapstructure:"logger"`
	Server   ServerConfig   `mapstructure:"server"`
}

func New() IConfig {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	if err := viper.ReadInConfig(); err != nil {
		panic("failed to read config file")
	}

	var instance config
	if err := viper.Unmarshal(&instance); err != nil {
		panic("failed to decode config")
	}

	return &instance
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

	return nil, exceptions.ErrInvalidClientName
}
