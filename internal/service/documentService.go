package service

import (
	"kz.nitec.digidocs.pcr/internal/config"
	"kz.nitec.digidocs.pcr/internal/models"
	"kz.nitec.digidocs.pcr/internal/repository"
)

type DocumentIfo struct {
	repo repository.PcrCertificate
	conf *config.Pcr
}

func NewDocumentService(repo repository.PcrCertificate, conf *config.Pcr) *DocumentIfo {
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
