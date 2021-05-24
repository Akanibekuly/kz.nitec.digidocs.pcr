package Service

import (
	"kz.nitec.digidocs.pcr/internal/config"
	"kz.nitec.digidocs.pcr/internal/repository"
)

type Services struct {
	PcrCertificateService ShepService
}

type ShepService interface {
	GetBySoap(interface{}) (interface{}, error)
}

type Deps struct {
	Repos *repository.Repositories
	ShepConfig *config.Shep
}

func NewServices(deps Deps) *Services{
	return &Services{
		PcrCertificateService: NewPcrCertificateService(deps.Repos.PcrCertificate, deps.ShepConfig),
	}
}