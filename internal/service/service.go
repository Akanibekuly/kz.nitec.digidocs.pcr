package service

import (
	"kz.nitec.digidocs.pcr/internal/config"
	"kz.nitec.digidocs.pcr/internal/models"
	"kz.nitec.digidocs.pcr/internal/repository"
)

type Services struct {
	PcrCertificateService ShepService
}


type ShepService interface {
	GetBySoap(*models.SoapRequest) (*models.SoapResponse, error)
	NewSoapRequest(*models.DocumentRequest) (*models.SoapRequest, error)
}

type Deps struct {
	Repos      *repository.Repositories
	ShepConfig *config.Shep
	PcrConfig  *config.Services
}

func NewServices(deps Deps) *Services {
	return &Services{
		PcrCertificateService: newPcrCertificateService(deps.Repos.ServiceRepo ,deps.ShepConfig, deps.PcrConfig),
	}
}
