package config

func GetApiVersion() string {
	return cfg.APIVersion
}

func GetJwtKey() string {
	return cfg.JwtKey
}

func GetDBConfig() *Database {
	if cfg.Database == nil {
		return DefaultDBConfig()
	}
	return cfg.Database
}

func DefaultDBConfig() *Database {
	return &Database{
		Host:     "localhost",
		Port:     "9090",
		Name:     "goroster",
		Username: "user",
		Password: "",
	}
}

func GetServerPort() string {
	return cfg.Port
}
