package repository

import (
	"database/sql"
	"fmt"
	"kz.nitec.digidocs.pcr/pkg/utils"
)

func NewPostgresDB(cfg *utils.DBConf) (*sql.DB, error) {
	var db *sql.DB
	var err error
	dbURI := fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Port,
		cfg.Host,
		cfg.Username,
		cfg.Password,
		cfg.DBName,
	)

	db, err = sql.Open(cfg.Dialect, dbURI)
	if err != nil {
		return nil, err
	}

	return db, nil
}
