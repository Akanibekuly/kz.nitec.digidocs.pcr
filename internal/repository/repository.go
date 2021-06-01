package repository

import (
	"database/sql"
	"kz.nitec.digidocs.pcr/internal/models"
)

type ServiceRepo interface {
	GetServiceInfoByCode(code string) (*models.Service, error)
	GetDocInfoByCode(code string) (*models.Document, error)
	GetServiceIdByCode(code string) (string, error)
	GetServiceUrlByCode(code string) (string, error)
}

type Repositories struct {
	ServiceRepo
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		NewServiceRepository(db),
	}
}
