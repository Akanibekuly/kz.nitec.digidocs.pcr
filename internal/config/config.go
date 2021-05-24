package config

type MainConfig struct {
	App  *AppConf
	DB   *DBConf
	Shep *Shep
	AppCode string
}

type (
	// AppConf represents ...
	AppConf struct {
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

	Shep struct {
		ShepEndpoint   string
		SenderLogin    string
		SenderPassword string
		ShepLogin      string
		ShepPassword   string
	}
)

func GetConfig() *MainConfig {
	return &MainConfig{
		App: &AppConf{
			Mode: getVarEnvAsStr("APP__MODE", "debug"),
			Port: getVarEnvAsStr("APP__PORT", "8080"),
		},
		DB: &DBConf{
			Dialect:  getVarEnvAsStr("DB_DIALECT", ""),
			Host:     getVarEnvAsStr("DB_HOST", ""),
			Port:     getVarEnvAsInt("DB_PORT", 0),
			Username: getVarEnvAsStr("DB_LOGIN", ""),
			Password: getVarEnvAsStr("DB_PASSWORD", ""),
			DBName:   getVarEnvAsStr("DB_NAME", ""),
		},
		Shep: &Shep{
			ShepEndpoint:   getVarEnvAsStr("SHEP_ENDPOINT", ""),
			SenderLogin:    getVarEnvAsStr("SENDER_LOGIN", ""),
			SenderPassword: getVarEnvAsStr("SENDER_PASSWORD", ""),
			ShepLogin:      getVarEnvAsStr("SHEP_LOGIN", ""),
			ShepPassword:   getVarEnvAsStr("SHEP_PASSWORD", ""),
		},
		AppCode: getVarEnvAsStr("APP_CODE","PCR_CERTIFICATE"),
	}
}
