package config

// TODO: IDP problem with auth. How we can auth-ze request? Current ios/android app don't send any auth-token/sso

type MainConfig struct {
	App *AppConf
	DB *DBConf
	// TODO: add if need more configs
}

type (
	// AppConf represents ...
	AppConf struct {
		// TODO: add if need more configs
		Mode string
		Port string
	}

	// DBConf represents ...
	DBConf struct {
		Dialect  string
		Host     string
		Port     int
		Username string
		Password string
		DBName   string
	}
)

func GetConfig() *MainConfig {
	return &MainConfig{
		App: &AppConf{
			Mode: getVarEnvAsStr("APP__MODE", "debug"),
			Port: getVarEnvAsStr("APP__PORT", "8080"),
		},
		DB: &DBConf{
			Dialect:  getVarEnvAsStr("", ""),
			Host:     getVarEnvAsStr("", ""),
			Port:     getVarEnvAsInt("", 0),
			Username: getVarEnvAsStr("", ""),
			Password: getVarEnvAsStr("", ""),
			DBName:   getVarEnvAsStr("", ""),
		},
	}
}


