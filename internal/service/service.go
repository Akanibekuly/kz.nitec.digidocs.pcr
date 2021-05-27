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
	BuildDocumentResponse(doc *models.Document, soap *models.SoapResponse) (*models.IssuedDigiDoc, error)
}

type ShepService interface {
	GetBySoap(request *models.SoapRequest, url string) (*models.SoapResponse, error)
	NewSoapRequest(*models.DocumentRequest, string) *models.SoapRequest
}

type Deps struct {
	Repos      *repository.Repositories
	ShepConfig *config.Shep
	PcrConfig  *config.Services
}

func NewServices(deps Deps) *Services {
	return &Services{
		PcrCertificateService: newPcrCertificateService(deps.ShepConfig, deps.PcrConfig.PcrCertificateCode),
		DocumentService:       newDocumentService(deps.Repos.PcrCertificate, deps.PcrConfig),
	}
}
