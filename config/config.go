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
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

// LoadConfig loads the configuration values from the config file
func LoadConfig() error {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return err
	}

	return nil
}
