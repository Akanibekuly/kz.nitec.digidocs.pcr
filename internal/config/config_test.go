// +build unit

package config

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
)

//# App configuration
//APP__MODE=debug
//APP__PORT=:8080
//
//# SHEP configurations
//SHEP_LOGIN=HDD1sQco26feoPeydupW
//SHEP_PASSWORD=83297abe28d0a43b204c783af341430a
//
//SHEP_ENDPOINT=http://192.168.145.133:9012/bip-sync/
//SENDER_LOGIN=mgov
//SENDER_PASSWORD=ff5d1ce5f9108f2581bedbf24d70af0e107ba21a1975f6fbb90c957223c05d06
//
//SHEP_RETRY_COUNT=1
//
//# DB configs
//DB_DIALECT=postgres
//DB_URI=192.168.217.107
//DB_PORT=5432
//DB_LOGIN=mgov
//DB_PASSWORD=mgov
//DB_NAME=digilocker
//
//#PCR
//PCR_CODE=PCR_CERTIFICATE
//PCR_NAME=PcrCertificate

func TestGetConfigAllOK(t *testing.T) {
	envs := map[string]string{
		"APP__MODE":        "debug",
		"APP__PORT":        ":8080",
		"SHEP_LOGIN":       "login",
		"SHEP_PASSWORD":    "password",
		"SHEP_ENDPOINT":    "http://127.0.0.1:9012/bip-sync/",
		"SENDER_LOGIN":     "mgov",
		"SENDER_PASSWORD":  "password",
		"SHEP_RETRY_COUNT": "1",
		"DB_DIALECT":       "postgres",
		"DB_URI":           "192.168.127.1",
		"DB_PORT":          "8080",
		"DB_LOGIN":         "login",
		"DB_PASSWORD":      "login",
		"DB_NAME":          "digilocker",
		"PCR_CODE":         "PCR_CERTIFICATE",
		"PCR_NAME":         "PcrCertificate",
	}

	for k, v := range envs {
		os.Setenv(k, v)
	}

	expectedResult := &MainConfig{
		App: &AppConf{
			Mode: "debug",
			Port: ":8080",
		},
		Shep: &Shep{
			SenderLogin:    "mgov",
			SenderPassword: "password",
			ShepLogin:      "login",
			ShepPassword:   "password",
			ShepRetryCount: 1,
		},
		DB: &DBConf{
			Dialect:  "postgres",
			Host:     "192.168.127.1",
			Port:     8080,
			Username: "login",
			Password: "login",
			DBName:   "digilocker",
		},
		Services: &Services{
			PcrCertificateCode:        "PCR_CERTIFICATE",
			PcrCertificateDocInfoCode: "PcrCertificate",
		},
	}

	result, err := GetConfig()

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(result, expectedResult)

	for k, _ := range envs {
		os.Unsetenv(k)
	}
}

func TestGetConfigEmptyEnv(t *testing.T) {
	result, err := GetConfig()
	assert := assert.New(t)
	assert.Nil(result)
	assert.Equal(err, fmt.Errorf("Env with key APP__MODE doesn't exists "))
}

func TestGetConfigDbPortIsNotInt(t *testing.T) {
	envs := map[string]string{
		"APP__MODE":        "debug",
		"APP__PORT":        ":8080",
		"SHEP_LOGIN":       "login",
		"SHEP_PASSWORD":    "password",
		"SHEP_ENDPOINT":    "http://127.0.0.1:9012/bip-sync/",
		"SENDER_LOGIN":     "mgov",
		"SENDER_PASSWORD":  "password",
		"SHEP_RETRY_COUNT": "1",
		"DB_DIALECT":       "postgres",
		"DB_URI":           "192.168.127.1",
		"DB_PORT":          "something",
		"DB_LOGIN":         "login",
		"DB_PASSWORD":      "login",
		"DB_NAME":          "digilocker",
		"PCR_CODE":         "PCR_CERTIFICATE",
		"PCR_NAME":         "PcrCertificate",
	}

	for k, v := range envs {
		os.Setenv(k, v)
	}


	result, err := GetConfig()

	assert := assert.New(t)
	assert.Nil(result)
	assert.Equal(err,&strconv.NumError{
		Func: "Atoi",
		Num: "something",
		Err: fmt.Errorf("invalid syntax"),
	})

	for k, _ := range envs {
		os.Unsetenv(k)
	}
}

func TestGetConfigShepRetryCountIsNotInt(t *testing.T) {
	envs := map[string]string{
		"APP__MODE":        "debug",
		"APP__PORT":        ":8080",
		"SHEP_LOGIN":       "login",
		"SHEP_PASSWORD":    "password",
		"SHEP_ENDPOINT":    "http://127.0.0.1:9012/bip-sync/",
		"SENDER_LOGIN":     "mgov",
		"SENDER_PASSWORD":  "password",
		"SHEP_RETRY_COUNT": "something",
		"DB_DIALECT":       "postgres",
		"DB_URI":           "192.168.127.1",
		"DB_PORT":          "8080",
		"DB_LOGIN":         "login",
		"DB_PASSWORD":      "login",
		"DB_NAME":          "digilocker",
		"PCR_CODE":         "PCR_CERTIFICATE",
		"PCR_NAME":         "PcrCertificate",
	}

	for k, v := range envs {
		os.Setenv(k, v)
	}


	result, err := GetConfig()

	assert := assert.New(t)
	assert.Nil(result)
	assert.Equal(err,&strconv.NumError{
		Func: "Atoi",
		Num: "something",
		Err: fmt.Errorf("invalid syntax"),
	})

	for k, _ := range envs {
		os.Unsetenv(k)
	}
}
