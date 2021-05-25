package repository

import (
	"database/sql"
	"kz.nitec.digidocs.pcr/internal/models"
)

type PcrCertificate interface {
	GetServiceInfoByCode(code string) (*models.Service, error)
	GetDocInfoByCode(code string) (*models.Document, error)
}

type Repositories struct {
	PcrCertificate
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		NewPcrRepository(db),
	}
}
