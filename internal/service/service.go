package service

import (
	"kz.nitec.digidocs.pcr/internal/models"
	"kz.nitec.digidocs.pcr/internal/repository"
	"kz.nitec.digidocs.pcr/pkg/utils"
)

type Services struct {
	PcrCertificateService ShepService
	DocumentService       DocumentService
}

type DocumentService interface {
	GetServiceInfoByCode() (*models.Service, error)
	GetDocInfoByCode() (*models.Document, error)
	BuildDocumentResponse(doc *models.Document,soap *models.SoapResponse) (*models.IssuedDigiDoc,error)
}

type ShepService interface {
	GetBySoap(request *models.SoapRequest, url string) (*models.SoapResponse, error)
	NewSoapRequest(*models.DocumentRequest, string) *models.SoapRequest
}

type Deps struct {
	Repos      *repository.Repositories
	ShepConfig *utils.Shep
	PcrConfig  *utils.Pcr
}

func NewServices(deps Deps) *Services {
	return &Services{
		PcrCertificateService: newPcrCertificateService(deps.ShepConfig, deps.PcrConfig.Code),
		DocumentService:       newDocumentService(deps.Repos.PcrCertificate, deps.PcrConfig),
	}
}
