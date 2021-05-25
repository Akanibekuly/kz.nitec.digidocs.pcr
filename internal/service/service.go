package service

import (
	"kz.nitec.digidocs.pcr/internal/config"
	"kz.nitec.digidocs.pcr/internal/models"
	"kz.nitec.digidocs.pcr/internal/repository"
)

type Services struct {
	PcrCertificateService ShepService
	DocumentService DocumentService
}

type DocumentService interface {
	GetServiceInfoByCode(code string) (*models.Service, error)
	GetDocInfoByCode(code string) (*models.Document, error)
}

type ShepService interface {
	GetBySoap(interface{}) (interface{}, error)
}

type Deps struct {
	Repos      *repository.Repositories
	ShepConfig *config.Shep
	Code       string
}

func NewServices(deps Deps) *Services {
	return &Services{
		PcrCertificateService: newPcrCertificateService(deps.Repos.PcrCertificate, deps.ShepConfig, deps.Code),
		DocumentService:  deps.Repos.PcrCertificate,
	}
}
