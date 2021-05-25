package service

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"kz.nitec.digidocs.pcr/internal/config"
	"kz.nitec.digidocs.pcr/internal/models"
	"kz.nitec.digidocs.pcr/internal/repository"
	"log"
	"net/http"
	"time"
)

const (
	ENVELOPE             = "Envelope"
	ENVELOP_SCHEMA       = "http://schemas.xmlsoap.org/soap/envelope/"
	SEND_MESSAGE_XMLNS   = "http://bip.bee.kz/SyncChannel/v10/Types"
	COVID_RESPONSE_XLMNS = "http://api.nce.kz/SyncChannel/v1/Types/CovidResponse"
	DIGILOCKER_XLMNS     = "http://digilocker.gov.kz/documentResponse/type/pcrcert"
	COVID_REQUEST_XLMNS  = "http://api.nce.kz/SyncChannel/v1/Types/CovidRequest"
	XSI_XMLNS_SCEMA      = "http://www.w3.org/2001/XMLSchema-instance"
	COVID_REQUEST_TYPE   = "ns6:CovidRequest"
)

type PcrCertificateService struct {
	repo repository.PcrCertificate
	conf *config.Shep
	code string
}

func newPcrCertificateService(repo repository.PcrCertificate, conf *config.Shep, code string) *PcrCertificateService {
	return &PcrCertificateService{
		repo, conf, code,
	}
}

func (pcr *PcrCertificateService) GetBySoap(request interface{}) (interface{}, error) {
	if request == nil {
		return nil, fmt.Errorf("Wrong request type %T ", request)
	}
	var docRequest *models.DocumentRequest
	switch request.(type) {
	case *models.DocumentRequest:
		docRequest = request.(*models.DocumentRequest)
	default:
		return nil, fmt.Errorf("Wrong request type %T ", request)
	}

	service := docRequest.Services["PCR_CERTIFICATE"]
	envelope := models.EnvelopeRequest{
		XMLName: xml.Name{Local: ENVELOPE},
		Text:    "",
		Xmlns:   ENVELOP_SCHEMA,
		Body: &models.BodyRequest{
			Text: "",
			SendMessage: &models.SendMessageRequest{
				Text: "",
				Ns2:  SEND_MESSAGE_XMLNS,
				Ns3:  COVID_RESPONSE_XLMNS,
				Ns4:  DIGILOCKER_XLMNS,
				Req: &models.Request{
					Text: "",
					ReqInfo: &models.RequestInfo{
						MessageId:   uuid.New().String(),
						MessageDate: time.Now().Format("2006-01-02T15:04:05Z07:00"),
						ServiceId:   service.ServiceId,
						Sender: &models.SenderCred{
							SenderId: pcr.conf.SenderLogin,
							Password: pcr.conf.SenderPassword,
						},
					},
					ReqData: &models.RequestData{
						Text: "",
						Data: &models.Data{
							Ns6:      COVID_REQUEST_XLMNS,
							Xsi:      XSI_XMLNS_SCEMA,
							Type:     COVID_REQUEST_TYPE,
							Iin:      docRequest.Iin,
							Login:    pcr.conf.ShepLogin,
							Password: pcr.conf.ShepPassword,
						},
					},
				},
			},
		},
	}

	b, err := xml.Marshal(envelope)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, pcr.conf.ShepEndpoint, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("could not read response body")
		return nil, err
	}

	shepResponse := &models.EnvelopeResponse{}
	err = xml.Unmarshal(response, shepResponse)
	if err != nil {
		fmt.Println(err)
	}

	docResponse, err := pcr.buildDocumentResponse(shepResponse)
	if err != nil {
		return nil, err
	}

	return docResponse, nil
}

// fill Document response for json response struct
func (pcr *PcrCertificateService) buildDocumentResponse(shepResponse *models.EnvelopeResponse) (*models.DocResponse, error) {
	var response *models.DocResponse

	return response, nil
}
