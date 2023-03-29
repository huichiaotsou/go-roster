package config

import (
	"github.com/spf13/viper"
)

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

func GetApiVersion() string {
	return cfg.APIVersion
}

func GetJwtKey() string {
	return cfg.JwtKey
}

func GetDBConfig() *Database {
	return cfg.Database
}

func GetServerPort() string {
	return cfg.Port
}
