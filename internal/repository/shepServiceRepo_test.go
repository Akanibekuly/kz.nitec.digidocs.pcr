// +build unit

package repository

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestServiceRepository_GetServiceIdByCodeOk(t *testing.T) {
	db, mock := NewMock()
	repo := NewServiceRepository(db)

	rows := sqlmock.NewRows([]string{"service_id"})
	rows.AddRow("PersonPhoto")
	query := "SELECT service_id FROM .+"
	mock.ExpectQuery(query).WithArgs("PERSON_PHOTO").WillReturnRows(rows)

	serviceId, err := repo.GetServiceIdByCode("PERSON_PHOTO")
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(serviceId, "PersonPhoto")
}

func TestServiceRepository_GetServiceIdByCodeWithError(t *testing.T) {
	db, mock := NewMock()
	repo := NewServiceRepository(db)

	query := "SELECT service_id FROM .+"
	mock.ExpectQuery(query).WithArgs("PERSON_PHOTO").WillReturnError(sql.ErrNoRows)

	serviceId, err := repo.GetServiceIdByCode("PERSON_PHOTO")
	assert := assert.New(t)
	assert.Equal(err, sql.ErrNoRows)
	assert.Equal(serviceId, "")
}

func TestServiceRepository_GetServiceUrlByCodeOk(t *testing.T) {
	db, mock := NewMock()
	repo := NewServiceRepository(db)

	rows := sqlmock.NewRows([]string{"url"})
	rows.AddRow("url")
	query := "SELECT url FROM .+"
	mock.ExpectQuery(query).WithArgs("url").WillReturnRows(rows)

	serviceId, err := repo.GetServiceUrlByCode("url")
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(serviceId, "url")
}

func TestServiceRepository_GetServiceUrlByCodeWithError(t *testing.T) {
	db, mock := NewMock()
	repo := NewServiceRepository(db)

	query := "SELECT url FROM .+"
	mock.ExpectQuery(query).WithArgs("url").WillReturnError(sql.ErrNoRows)

	serviceId, err := repo.GetServiceUrlByCode("url")
	assert := assert.New(t)
	assert.Equal(err, sql.ErrNoRows)
	assert.Equal(serviceId, "")
}
