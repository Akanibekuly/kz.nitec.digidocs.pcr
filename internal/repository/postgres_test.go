package repository

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"kz.nitec.digidocs.pcr/internal/config"
	"testing"
)

func TestNewPostgresDBOk(t *testing.T) {
	conf:=&config.DBConf{
		"postgres",
		"127.0.0.1",
		5432,
		"login",
		"pass",
		"db",
	}

	_,err:=NewPostgresDB(conf)
	assert:=assert.New(t)
	assert.Nil(err)
}

func TestNewPostgresDBWithError(t *testing.T) {
	conf:=&config.DBConf{
		"error",
		"127.0.0.1",
		5432,
		"login",
		"pass",
		"db",
	}

	_,err:=NewPostgresDB(conf)
	assert:=assert.New(t)
	assert.Equal(fmt.Errorf("sql: unknown driver \"error\" (forgotten import?)"),err)
}