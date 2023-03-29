package config

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
	MaxopenConns int    `mapstructure:"max_open_connections"`
	MaxIdleConns int    `mapstructure:"max_idle_connections"`
}
