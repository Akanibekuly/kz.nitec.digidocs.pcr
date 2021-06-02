package repository

import (
	"database/sql"
	"kz.nitec.digidocs.pcr/internal/models"
	"kz.nitec.digidocs.pcr/pkg/logger"
)

type BuildServiceRepository struct {
	db *sql.DB
}

func NewBuildServiceRepsoitory(db *sql.DB) *BuildServiceRepository {
	return &BuildServiceRepository{db}
}

func (brepo *BuildServiceRepository) GetDocInfoByCode(code string) (*models.Document, error) {
	row := brepo.db.QueryRow("SELECT name_en, name_ru, name_kk FROM document_type WHERE code=$1", code)
	doc := &models.Document{Code: code}
	err := row.Scan(&doc.NameEn, &doc.NameRu, &doc.NameKK)
	if err != nil {
		return nil, logger.CreateMessageLog(err)
	}

	return doc, nil
}
