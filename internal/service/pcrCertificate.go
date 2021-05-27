package service

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"kz.nitec.digidocs.pcr/internal/models"
	"kz.nitec.digidocs.pcr/pkg/utils"
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
	conf *utils.Shep
	code string
}

func newPcrCertificateService(conf *utils.Shep, code string) *PcrCertificateService {
	return &PcrCertificateService{
		conf, code,
	}
}

func (pcr *PcrCertificateService) GetBySoap(soapRequest *models.SoapRequest, url string) (*models.SoapResponse, error) {
	b, err := xml.Marshal(soapRequest)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("could not read response body")
		return nil, err
	}

	shepResponse := &models.SoapResponse{}
	err = xml.Unmarshal(data, shepResponse)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return shepResponse, nil
}

func (pcr *PcrCertificateService) NewSoapRequest(request *models.DocumentRequest, serviceId string) *models.SoapRequest {
	return &models.SoapRequest{
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
						ServiceId:   serviceId,
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
							Iin:      request.Iin,
							Login:    pcr.conf.ShepLogin,
							Password: pcr.conf.ShepPassword,
						},
					},
				},
			},
		},
	}
}
