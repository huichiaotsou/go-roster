package config

import (
	"github.com/spf13/viper"
)

var (
	cfg *Config
)

// Config represents the configuration values
type Config struct {
	Database   *Database `mapstructure:"database"`
	Port       string    `mapstructure:"port"`
	Env        string    `mapstructure:"env"`
	APIVersion string    `mapstructure:"apiVersion"`
	JwtKey     string    `mapstructure:"jwtKey"`
}

type Database struct {
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	Name         string `mapstructure:"name"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Schema       string `mapstructure:"schema"`
	MaxopenConns string `mapstructure:"max_open_connections"`
	MaxIdleConns string `mapstructure:"max_idle_connections"`
}

// LoadConfig loads the configuration values from the config file
func LoadConfig() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err = viper.ReadInConfig(); err != nil {
		return err
	}

	if err = viper.Unmarshal(&cfg); err != nil {
		return err
	}

	return nil
}
