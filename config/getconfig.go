package config

func GetApiVersion() string {
	if cfg.APIVersion == "" {
		return "v1"
	}
	return cfg.APIVersion
}

func GetDBConfig() *Database {
	if cfg.Database == nil {
		return &Database{}
	}
	return cfg.Database
}

func GetServerPort() string {
	return cfg.Port
}
