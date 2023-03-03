package config

import (
	"fmt"
	"log"

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
}

type Database struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

// LoadConfig loads the configuration values from the config file
func LoadConfig() {
	if cfg != nil {
		log.Fatal("config exists")
		return
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Error reading config file: %v", err))
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(fmt.Sprintf("Error unmarshaling config file: %v", err))
	}
}
