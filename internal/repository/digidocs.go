package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"kz.nitec.digidocs.pcr/internal/models"
	"log"
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
		log.Println(err)
		return "", err
	}

	return result, nil
}

func (repo *ServiceRepository) GetServiceUrlByCode(code string) (string, error) {
	row := repo.db.QueryRow("SELECT url FROM service WHERE code = $1", code)
	var result string
	err := row.Scan(&result)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return result, nil
}

func (repo *ServiceRepository) GetServiceInfoByCode(code string) (*models.Service, error) {
	row := repo.db.QueryRow("SELECT service_id, url FROM service WHERE code = $1", code)
	service := &models.Service{
		Code: code,
	}
	err := row.Scan(&service.ServiceId, &service.URL)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return service, nil
}

func (repo *ServiceRepository) GetDocInfoByCode(code string) (*models.Document, error) {
	row := repo.db.QueryRow("SELECT name_en, name_ru, name_kk FROM document_type WHERE code=$1", code)
	doc := &models.Document{Code: code}
	err := row.Scan(&doc.NameEn, &doc.NameRu, &doc.NameKK)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
