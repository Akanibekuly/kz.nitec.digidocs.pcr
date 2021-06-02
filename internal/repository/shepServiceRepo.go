package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"kz.nitec.digidocs.pcr/pkg/logger"
)

type ServiceRepository struct {
	db *sql.DB
}

func NewServiceRepository(db *sql.DB) *ServiceRepository {
	return &ServiceRepository{db: db}
}

func (repo *ServiceRepository) GetServiceIdByCode(code string) (string, error) {
	row := repo.db.QueryRow("SELECT service_id FROM service WHERE code = $1", code)
	var result string
	err := row.Scan(&result)
	if err != nil {
		return "", logger.CreateMessageLog(err)
	}

	return result, nil
}

func (repo *ServiceRepository) GetServiceUrlByCode(code string) (string, error) {
	row := repo.db.QueryRow("SELECT url FROM service WHERE code = $1", code)
	var result string
	err := row.Scan(&result)
	if err != nil {
		return "", logger.CreateMessageLog(err)
	}

	return result, nil
}
