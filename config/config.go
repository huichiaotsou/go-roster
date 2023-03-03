package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Config represents the configuration values
type Config struct {
	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Name     string `mapstructure:"name"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	} `mapstructure:"database"`
	Port       int    `mapstructure:"port"`
	Env        string `mapstructure:"env"`
	APIVersion string `mapstructure:"apiVersion"`
}

var Cfg *Config

// LoadConfig loads the configuration values from the config file
func LoadConfig() {
	if Cfg != nil {
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

	err = viper.Unmarshal(&Cfg)
	if err != nil {
		panic(fmt.Sprintf("Error unmarshaling config file: %v", err))
	}

}
