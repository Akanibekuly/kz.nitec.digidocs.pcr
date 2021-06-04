package service

import (
	"fmt"
	"kz.nitec.digidocs.pcr/internal/config"
	"kz.nitec.digidocs.pcr/internal/models"
	"kz.nitec.digidocs.pcr/internal/repository"
	"kz.nitec.digidocs.pcr/pkg/logger"
	"time"
)

type BuildService struct {
	repo repository.BuildServiceRepo
	conf *config.Services
}

func newBuildService(repo repository.BuildServiceRepo, conf *config.Services) *BuildService {
	return &BuildService{repo: repo, conf: conf}
}

func (b *BuildService) BuildDocumentResponse(response *models.SoapResponse) (*models.IssuedDigiDoc, error) {
	covidResult, err := getCovidResult(response.Body.SendMessageResponse.Response.ResponseData.Data.Result.Covid)
	if err != nil {
		return nil, logger.CreateMessageLog(err)
	}

	document,err:=b.repo.GetDocInfoByCode(b.conf.PcrCertificateDocInfoCode)
	if err!=nil{
		return nil,logger.CreateMessageLog(err)
	}

	common := models.DocCommon{
		DocOwner: models.DocPerson{
			Iin:        covidResult.Patient.IIN,
			FirstName:  covidResult.Patient.FirstName,
			LastName:   covidResult.Patient.LastName,
			MiddleName: covidResult.Patient.MiddleName,
		},
		DocType: models.DocType{
			Code: document.Code,
			I18Text: models.I18Text{
				NameEn: document.NameEn,
				NameRu: document.NameRu,
				NameKk: document.NameKK,
			},
		},
		DocUri:  fmt.Sprintf("%s:%s", document.Code, covidResult.Key),
	}

	isResident := "Ия/Да"
	if isResident != "true" {
		isResident = "Жоқ/Нет"
	}
	isHasSymptomsCOVID := "Жоқ/Нет"
	if covidResult.HasSymptomsCOVID == "true" {
		isHasSymptomsCOVID = "Ия/Да"
	}
	isResearchResults := "Теріс/Отрицательный"
	if covidResult.ResearchResults == "true" {
		isResearchResults = "Оң/Положительный"
	}
	gender := "Ер/Мужской"
	if covidResult.Patient.Gender != "MAN_ENUM" {
		gender = "Әйел/Женский"
	}

	pcrCertificate := models.PcrCertificate{
		Key:              covidResult.Key,
		FirstName:        covidResult.Patient.FirstName,
		LastName:         covidResult.Patient.LastName,
		MiddleName:       covidResult.Patient.MiddleName,
		Iin:              covidResult.Patient.IIN,
		Adress:           covidResult.Patient.AddressOfActualResidence,
		Birthday:         covidResult.Patient.Birthday,
		PlaceOfStudy:     covidResult.Patient.PlaceOfStudyOrWork,
		ProtocolDate:     covidResult.ProtocolDate,
		CreateAt:         covidResult.CreatedAt,
		IsResident:       isResident,
		HasSymptomsCOVID: isHasSymptomsCOVID,
		ResearchResults:  isResearchResults,
		Gender:           gender,
	}

	digidoc := &models.IssuedDigiDoc{
		Common: common,
		Domain: pcrCertificate,
	}

	return digidoc, nil
}

func getCovidResult(results []models.CovidResult) (*models.CovidResult, error) {
	var result *models.CovidResult
	for _, v := range results {
		if result == nil {
			result = &v
		} else {
			t1, err := time.Parse("2006-01-02T15:04:05Z07:00", v.ProtocolDate)
			if err != nil {
				return nil, logger.CreateMessageLog(err)
			}
			t2, err := time.Parse("2006-01-02T15:04:05Z07:00", result.ProtocolDate)
			if err != nil {
				return nil, logger.CreateMessageLog(err)
			}

			if t1.After(t2) {
				result = &v
			}
		}
	}
	return result, nil
}
