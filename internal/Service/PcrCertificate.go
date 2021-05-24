package Service

import (
	"kz.nitec.digidocs.pcr/internal/config"
	"kz.nitec.digidocs.pcr/internal/repository"
)

type PcrCertificateService struct {
	repo repository.PcrCertificate
	conf *config.Shep
}

func NewPcrCertificateService(repo repository.PcrCertificate, conf *config.Shep) *PcrCertificateService{
	return &PcrCertificateService{
		repo: repo,
		conf: conf,
	}
}

func (pcr *PcrCertificateService) GetBySoap(request interface{}) (interface{},error){

	return nil,nil
}
