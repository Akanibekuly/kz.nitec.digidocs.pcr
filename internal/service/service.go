package service

import (
	"kz.nitec.digidocs.pcr/internal/config"
	"kz.nitec.digidocs.pcr/internal/models"
	"kz.nitec.digidocs.pcr/internal/repository"
)

type Services struct {
	PcrCertificateService ShepService
	DocumentService       DocumentService
}

type DocumentService interface {
	GetServiceInfoByCode() (*models.Service, error)
	GetDocInfoByCode() (*models.Document, error)
}

type ShepService interface {
	GetBySoap(request *models.SoapRequest, url string) (*models.SoapResponse, error)
	NewSoapRequest(string, *models.DocumentRequest) *models.SoapRequest
}

type Deps struct {
	Repos      *repository.Repositories
	ShepConfig *config.Shep
	PcrConfig  *config.Pcr
	Code       string
}

func NewServices(deps Deps) *Services {
	return &Services{
		PcrCertificateService: newPcrCertificateService(deps.Repos.PcrCertificate, deps.ShepConfig, deps.Code),
		DocumentService:       NewDocumentService(deps.Repos.PcrCertificate, deps.PcrConfig),
	}
}
