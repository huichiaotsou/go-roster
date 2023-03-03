package config

func GetApiVersion() string {
	if cfg.APIVersion == "" {
		return "v1"
	}
	return cfg.APIVersion
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
