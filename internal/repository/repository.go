package repository

import "database/sql"

type PcrCertificate interface {
	GetServiceInfoByCode(code string) (*Service, error)
	GetDocInfoByCode(code string) (*Document, error)
}

type Repositories struct {
	PcrCertificate
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		NewPcrRepository(db),
	}
}
