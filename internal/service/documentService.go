package service

import (
	"database/sql"
	"kz.nitec.digidocs.pcr/internal/repository"
)

func NewDocumentService(db *sql.DB) *repository.PcrRepository{
	return repository.NewPcrRepository(db)
}