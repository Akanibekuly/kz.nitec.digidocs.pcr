package models

type DocResponse struct {
	Found     bool
	Active    bool
	Documents []*DocumentResponseItem
}

type DocumentResponseItem struct {
	Valid            bool
	Number           string
	Title            string
	TitleRu          string
	TitleKk          string
	IssuedDate       string
	ExpirationDate   string
	SerializableData IssuedDigiDoc
	StorageID        string
}

type IssuedDigiDoc struct {
	Common DocCommon
	Domain PcrCertificate
}

type DocCommon struct {
	DocOwner          DocPerson
	DocType           DocType
	DocNumber         string
	DocIssuer         string
	DocIssuedDate     string
	DocExpirationDate string
	DocUri            string
}

type DocPerson struct {
	Iin       string
	FirstName string
	LastName  string
	MiddleName string
}

type DocType struct {
	Code string
	I18Text
}

type I18Text struct {
	NameKk string
	NameRu string
	NameEn string
}

type PcrCertificate struct {
	Key              string
	FirstName        string
	LastName         string
	MiddleName       string
	Iin              string
	Gender           string
	IsResident       string
	Adress           string
	Birthday         string
	PlaceOfStudy     string
	Phone            string
	HasSymptomsCOVID string
	ProtovolDate     string
	CreateAt         string
	ResearchResults  string
}
