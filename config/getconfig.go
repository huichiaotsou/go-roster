package config

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
