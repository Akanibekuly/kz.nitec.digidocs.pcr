package repository

import "database/sql"

type PcrRepository struct {
	db *sql.DB
}

func NewPcrRepository(db *sql.DB) *PcrRepository {
	return &PcrRepository{db: db}
}

type Service struct {
	Code      string
	ServiceId sql.NullString
	URL       sql.NullString
}

func (repo *PcrRepository) GetServiceInfoByCode(code string) (*Service, error) {
	row := repo.db.QueryRow("SELECT service_id, url FROM service WHERE code = ?", code)
	service := &Service{
		Code: code,
	}
	err := row.Scan(&service.ServiceId, &service.URL)
	if err != nil {
		return nil, err
	}

	return service, nil
}

type Document struct {
	Code   string
	NameKK sql.NullString
	NameRu sql.NullString
	NameEn sql.NullString
}

func (repo *PcrRepository) GetDocInfoByCode(code string) (*Document, error) {
	row := repo.db.QueryRow("SELECT name_en, name_ru, name_kk FROM document_type WHERE code=?", code)
	doc := &Document{Code: code}
	err := row.Scan(&doc.NameEn, &doc.NameRu, &doc.NameKK)
	if err != nil {
		return nil, err
	}

	return doc, nil
}