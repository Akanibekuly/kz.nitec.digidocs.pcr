package config

import (
	"fmt"
	"os"
	"strconv"
)

type MainConfig struct {
	App      *AppConf
	DB       *DBConf
	Shep     *Shep
	Services *Services
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
	Services struct {
		PcrCertificateCode        string
		PcrCertificateDocInfoCode string
		ENVELOPE                  string
		ENVELOP_SCHEMA            string
		SEND_MESSAGE_XMLNS        string
		COVID_RESPONSE_XLMNS      string
		DIGILOCKER_XLMNS          string
		XSI_XMLNS_SCEMA           string
		COVID_REQUEST_TYPE        string
		SENDER_CODE               string
		COVID_REQUEST_XLMNS       string
	}
)

func GetConfig() (*MainConfig, error) {
	envs := []string{
		"APP__MODE", "APP__PORT",
		"DB_DIALECT", "DB_URI", "DB_PORT", "DB_LOGIN", "DB_PASSWORD", "DB_NAME",
		"SENDER_LOGIN", "SENDER_PASSWORD", "SHEP_LOGIN", "SHEP_PASSWORD", "SHEP_RETRY_COUNT",
		"PCR_CODE", "PCR_NAME",
		"ENVELOPE", "ENVELOP_SCHEMA", "SEND_MESSAGE_XMLNS", "COVID_RESPONSE_XLMNS", "DIGILOCKER_XLMNS", "XSI_XMLNS_SCEMA", "COVID_REQUEST_TYPE", "SENDER_CODE", "COVID_REQUEST_XLMNS",
	}

	for _, key := range envs {
		if val, exists := os.LookupEnv(key); !exists || val == "" {
			return nil, fmt.Errorf("Env with key %s doesn't exists ", key)
		}
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}

	retryCount, err := strconv.Atoi(os.Getenv("SHEP_RETRY_COUNT"))
	if err != nil {
		return nil, err
	}

	return &MainConfig{
		App: &AppConf{
			Mode: os.Getenv("APP__MODE"),
			Port: os.Getenv("APP__PORT"),
		},
		DB: &DBConf{
			Dialect:  os.Getenv("DB_DIALECT"),
			Host:     os.Getenv("DB_URI"),
			Port:     port,
			Username: os.Getenv("DB_LOGIN"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
		},
		Shep: &Shep{
			SenderLogin:    os.Getenv("SENDER_LOGIN"),
			SenderPassword: os.Getenv("SENDER_PASSWORD"),
			ShepLogin:      os.Getenv("SHEP_LOGIN"),
			ShepPassword:   os.Getenv("SHEP_PASSWORD"),
			ShepRetryCount: retryCount,
		},
		Services: &Services{
			PcrCertificateCode:        os.Getenv("PCR_CODE"),
			PcrCertificateDocInfoCode: os.Getenv("PCR_NAME"),
			ENVELOPE:                  os.Getenv("ENVELOPE"),
			ENVELOP_SCHEMA:            os.Getenv("ENVELOP_SCHEMA"),
			SEND_MESSAGE_XMLNS:        os.Getenv("SEND_MESSAGE_XMLNS"),
			COVID_RESPONSE_XLMNS:      os.Getenv("COVID_RESPONSE_XLMNS"),
			DIGILOCKER_XLMNS:          os.Getenv("DIGILOCKER_XLMNS"),
			XSI_XMLNS_SCEMA:           os.Getenv("XSI_XMLNS_SCEMA"),
			COVID_REQUEST_TYPE:        os.Getenv("COVID_REQUEST_TYPE"),
			SENDER_CODE:               os.Getenv("SENDER_CODE"),
			COVID_REQUEST_XLMNS:       os.Getenv("COVID_REQUEST_XLMNS"),
		},
	}, nil
}
