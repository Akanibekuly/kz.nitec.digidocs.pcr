package repository

import "database/sql"

type PcrRepository struct {
	db *sql.DB
}

func PcrRepositoryInit(db *sql.DB) *PcrRepository {
	return &PcrRepository{db: db}
}

func (d *PcrRepository) Mock() {}
