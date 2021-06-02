package repository

import (
	"database/sql"
	"kz.nitec.digidocs.pcr/internal/models"
)

type ServiceRepo interface {
	GetServiceIdByCode(code string) (string, error)
	GetServiceUrlByCode(code string) (string, error)
}

type BuildServiceRepo interface {
	GetDocInfoByCode(code string) (*models.Document, error)
}

type Repositories struct {
	ServiceRepo
	BuildServiceRepo
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		NewServiceRepository(db),
		NewBuildServiceRepsoitory(db),
	}
}
