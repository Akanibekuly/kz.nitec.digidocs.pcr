package repository

import (
	"database/sql"
	"fmt"
	"kz.nitec.digidocs.pcr/internal/config"
	"kz.nitec.digidocs.pcr/pkg/logger"
)

func NewPostgresDB(cfg *config.DBConf) (*sql.DB, error) {
	dbURI := fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Port,
		cfg.Host,
		cfg.Username,
		cfg.Password,
		cfg.DBName,
	)

	db, err := sql.Open(cfg.Dialect, dbURI)
	if err != nil {
		return nil, logger.CreateMessageLog(err)
	}

	return db, nil
}
