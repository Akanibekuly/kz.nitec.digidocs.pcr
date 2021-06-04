package repository

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"kz.nitec.digidocs.pcr/internal/models"
	"testing"
)

func NewMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock){
	db,mock,err:=sqlmock.New()
	if err!=nil{
		t.Errorf("%s",err)
		return nil,nil
	}
	return db,mock
}


func TestBuildServiceRepository_GetDocInfoByCode(t *testing.T) {
	assert:=assert.New(t)
	db,mock:=NewMock(t)
	rows:=mock.NewRows([]string{"name_en","name_ru", "name_kk"}).AddRow("english","орысша","қазақша")
	mock.ExpectQuery("SELECT name_en, name_ru, name_kk FROM document_type.+").WithArgs("args").WillReturnRows(rows)

	repo:=NewBuildServiceRepsoitory(db)
	docType,err:=repo.GetDocInfoByCode("args")
	assert.Nil(err)
	assert.Equal(&models.Document{
		Code: "args",
		NameKK: "қазақша",
		NameRu: "орысша",
		NameEn: "english",
	}, docType)
}

func TestBuildServiceRepository_GetDocInfoByCodeWithError(t *testing.T) {
	assert:=assert.New(t)
	db,mock:=NewMock(t)

	mock.ExpectQuery("SELECT name_en, name_ru, name_kk FROM document_type.+").WithArgs("args").WillReturnError(sql.ErrNoRows)

	repo:=NewBuildServiceRepsoitory(db)
	docType,err:=repo.GetDocInfoByCode("args")
	assert.Equal(fmt.Errorf("in function:GetDocInfoByCode file:buildServiceRepo.go line:22 message:sql: no rows in result set"),err)
	assert.Nil(docType)
}
