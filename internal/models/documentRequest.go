package models

type DocumentRequest struct {
	Iin             string
	Services        map[string]ServiceDTO
	DocumentTypeDto DocumentTypeDto `json:"documentType"`
}

type ServiceDTO struct {
	ID        float64
	Code      string
	Name      string
	Url       string
	ServiceId string
}

type DocumentTypeDto struct {
	ID     float64
	Code   string
	NameRu string
	NameKk string
	NameEn string
}
