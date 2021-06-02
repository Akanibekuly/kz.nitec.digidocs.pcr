package service

import (
	"kz.nitec.digidocs.pcr/internal/config"
	"kz.nitec.digidocs.pcr/internal/models"
	"kz.nitec.digidocs.pcr/internal/repository"
)

type Services struct {
	PcrCertificateService ShepService
	BuildService          Builder
}

type ShepService interface {
	GetBySoap(*models.SoapRequest) (*models.SoapResponse, error)
	NewSoapRequest(*models.DocumentRequest) (*models.SoapRequest, error)
}

type Builder interface {
	BuildDocumentResponse(response *models.SoapResponse) (*models.IssuedDigiDoc, error)
}

type Deps struct {
	Repos      *repository.Repositories
	ShepConfig *config.Shep
	PcrConfig  *config.Services
}

func NewServices(deps Deps) *Services {
	return &Services{
		PcrCertificateService: newPcrCertificateService(deps.Repos.ServiceRepo, deps.ShepConfig, deps.PcrConfig),
		BuildService:          newBuildService(deps.Repos.BuildServiceRepo),
	}
}
