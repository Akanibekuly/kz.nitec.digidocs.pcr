package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"kz.nitec.digidocs.pcr/internal/models"
)

type PcrRepository struct {
	db *sql.DB
}

func NewPcrRepository(db *sql.DB) *PcrRepository {
	return &PcrRepository{db: db}
}

func (repo *PcrRepository) GetServiceInfoByCode(code string) (*models.Service, error) {
	row := repo.db.QueryRow("SELECT service_id, url FROM service WHERE code = ?", code)
	service := &models.Service{
		Code: code,
	}
	err := row.Scan(&service.ServiceId, &service.URL)
	if err != nil {
		return nil, err
	}

	return service, nil
}

func (repo *PcrRepository) GetDocInfoByCode(code string) (*models.Document, error) {
	row := repo.db.QueryRow("SELECT name_en, name_ru, name_kk FROM document_type WHERE code=?", code)
	doc := &models.Document{Code: code}
	err := row.Scan(&doc.NameEn, &doc.NameRu, &doc.NameKK)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
