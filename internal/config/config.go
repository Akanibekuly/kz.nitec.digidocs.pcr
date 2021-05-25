package config

import (
	"fmt"
	"os"
	"strconv"
)

type MainConfig struct {
	App  *AppConf
	DB   *DBConf
	Shep *Shep
	Pcr  *Pcr
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
		SenderLogin    string
		SenderPassword string
		ShepLogin      string
		ShepPassword   string
		ShepRetryCount int
	}

	Pcr struct {
		Code string
		Name string
	}
)

func checkConfig() error {
	envs := []string{
		"APP__MODE", "APP__PORT",
		"DB_DIALECT", "DB_URI", "DB_PORT", "DB_LOGIN", "DB_PASSWORD", "DB_NAME",
		"SENDER_LOGIN", "SENDER_PASSWORD", "SHEP_LOGIN", "SHEP_PASSWORD", "SHEP_RETRY_COUNT",
		"PCR_CODE", "PCR_NAME",
	}

	for _, val := range envs {
		if key, exists := os.LookupEnv(val); !exists || key == "" {
			return fmt.Errorf("Env with key %s doesn't exists ", val)
		}
	}

	return nil
}

func GetConfig() (*MainConfig, error) {
	if err := checkConfig(); err != nil {
		return nil, err
	}
	return &MainConfig{
		App: &AppConf{
			Mode: getVarEnvAsStr("APP__MODE", "debug"),
			Port: getVarEnvAsStr("APP__PORT", "8080"),
		},
		DB: &DBConf{
			Dialect:  getVarEnvAsStr("DB_DIALECT", ""),
			Host:     getVarEnvAsStr("DB_URI", ""),
			Port:     getVarEnvAsInt("DB_PORT", 0),
			Username: getVarEnvAsStr("DB_LOGIN", ""),
			Password: getVarEnvAsStr("DB_PASSWORD", ""),
			DBName:   getVarEnvAsStr("DB_NAME", ""),
		},
		Shep: &Shep{
			SenderLogin:    getVarEnvAsStr("SENDER_LOGIN", ""),
			SenderPassword: getVarEnvAsStr("SENDER_PASSWORD", ""),
			ShepLogin:      getVarEnvAsStr("SHEP_LOGIN", ""),
			ShepPassword:   getVarEnvAsStr("SHEP_PASSWORD", ""),
			ShepRetryCount: getVarEnvAsInt("SHEP_RETRY_COUNT", 0),
		},
		Pcr: &Pcr{
			Code: getVarEnvAsStr("PCR_CODE", ""),
			Name: getVarEnvAsStr("PCR_NAME", ""),
		},
	}, nil
}

func getVarEnvAsStr(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getVarEnvAsInt(key string, defaultVal int) int {
	valueStr := getVarEnvAsStr(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}
