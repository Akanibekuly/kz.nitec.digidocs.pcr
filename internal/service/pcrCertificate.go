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

type PcrCertificateService struct {
	repo        repository.ServiceRepo
	conf        *config.Shep
	serviceConf *config.Services
}

func newPcrCertificateService(repo repository.ServiceRepo, conf *config.Shep, serviceConf *config.Services) *PcrCertificateService {
	return &PcrCertificateService{
		repo, conf, serviceConf,
	}
}

func (pcr *PcrCertificateService) GetBySoap(soapRequest *models.SoapRequest) (*models.SoapResponse, error) {
	url, err := pcr.repo.GetServiceUrlByCode(pcr.serviceConf.PcrCertificateCode)
	if err != nil {
		return nil, err
	}
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

func (pcr *PcrCertificateService) NewSoapRequest(soapRequest *models.DocumentRequest) (*models.SoapRequest, error) {
	serviceId, err := pcr.repo.GetServiceIdByCode(pcr.serviceConf.PcrCertificateCode)
	if err != nil {
		return nil, err
	}
	return &models.SoapRequest{
		XMLName: xml.Name{Local: pcr.serviceConf.ENVELOPE},
		Text:    "",
		Xmlns:   pcr.serviceConf.ENVELOP_SCHEMA,
		Body: &models.BodyRequest{
			Text: "",
			SendMessage: &models.SendMessageRequest{
				Text: "",
				Ns2:  pcr.serviceConf.SEND_MESSAGE_XMLNS,
				Ns3:  pcr.serviceConf.COVID_RESPONSE_XLMNS,
				Ns4:  pcr.serviceConf.DIGILOCKER_XLMNS,
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
							Ns6:      pcr.serviceConf.COVID_REQUEST_XLMNS,
							Xsi:      pcr.serviceConf.XSI_XMLNS_SCEMA,
							Type:     pcr.serviceConf.COVID_REQUEST_TYPE,
							Iin:      soapRequest.Iin,
							Login:    pcr.conf.ShepLogin,
							Password: pcr.conf.ShepPassword,
						},
					},
				},
			},
		},
	}, nil
}
