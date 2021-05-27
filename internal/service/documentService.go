package service

import (
	"kz.nitec.digidocs.pcr/internal/models"
	"kz.nitec.digidocs.pcr/internal/repository"
	"kz.nitec.digidocs.pcr/pkg/utils"
	"os"
	"time"
)

type DocumentIfo struct {
	repo repository.PcrCertificate
	conf *utils.Pcr
}

func newDocumentService(repo repository.PcrCertificate, conf *utils.Pcr) *DocumentIfo {
	return &DocumentIfo{
		repo, conf,
	}
}

func (dc *DocumentIfo) GetServiceInfoByCode() (*models.Service, error) {
	return dc.repo.GetServiceInfoByCode(dc.conf.Code)
}

func (dc *DocumentIfo) GetDocInfoByCode() (*models.Document, error) {
	return dc.repo.GetDocInfoByCode(dc.conf.Name)
}

func (dc *DocumentIfo) BuildDocumentResponse(doc *models.Document,soap *models.SoapResponse) (*models.IssuedDigiDoc,error){
	docResponse:=&models.IssuedDigiDoc{}
	covidResult,err:=getCovidResult(soap)
	if err!=nil{
		return nil,err
	}
	if covidResult==nil{
		return docResponse, nil
	}

	docOwner:=models.DocPerson{
		Iin: covidResult.Patient.IIN,
		FirstName: covidResult.Patient.FirstName,
		LastName: covidResult.Patient.LastName,
		MiddleName: covidResult.Patient.MiddleName,
	}
	docType:=models.DocType{
		Code: doc.Code,
		I18Text: models.I18Text{
			NameKk: doc.NameKK,
			NameRu: doc.NameRu,
			NameEn: doc.NameEn,
		},
	}
	common:=models.DocCommon{
		DocOwner: docOwner,
		DocType: docType,
		DocExpirationDate: covidResult.CreatedAt,
		DocIssuedDate: covidResult.ProtocolDate,
		DocNumber: covidResult.Key,
		DocUri: os.Getenv("PCR_CODE")+"+"+covidResult.Key,
	}

	domain:=models.PcrCertificate{
		Key: covidResult.Key,
		FirstName: covidResult.Patient.FirstName,
		LastName: covidResult.Patient.LastName,
		MiddleName: covidResult.Patient.MiddleName,
		Iin: covidResult.Patient.IIN,
		Gender: covidResult.Patient.Gender,

	}

	docResponse.Common=common
	docResponse.Domain=domain
	return docResponse,nil
}

func getCovidResult(soap *models.SoapResponse) (*models.CovidResult,error) {
	var covidResult *models.CovidResult
	covidResults:=soap.Body.SendMessageResponse.Response.ResponseData.Data.Result.Covid
	if len(covidResults)==0{
		return covidResult,nil
	}

	for _,val:=range covidResults{
		if covidResult==nil{
			covidResult=&val
		} else{
			t1,err:=time.Parse("2006-01-02T15:04:05Z07:00",val.ProtocolDate)
			if err!=nil{
				return nil,err
			}
			t2,err:=time.Parse("2006-01-02T15:04:05Z07:00",covidResult.ProtocolDate)
			if err!=nil{
				return nil,err
			}
			if t1.Before(t2){
				covidResult=&val
			}
		}
	}
	return covidResult,nil
}
