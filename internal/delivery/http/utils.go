package http

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	models2 "kz.nitec.digidocs.pcr/internal/models"
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

func (a *App) SendMessage(docRequest *models2.DocumentRequest) (*models2.EnvelopeResponse, error) {
	service := docRequest.Services["PCR_CERTIFICATE"]
	envelope := models2.EnvelopeRequest{
		XMLName: xml.Name{Local: ENVELOPE},
		Text:    "",
		Xmlns:   ENVELOP_SCHEMA,
		Body: &models2.BodyRequest{
			Text: "",
			SendMessage: &models2.SendMessageRequest{
				Text: "",
				Ns2:  SEND_MESSAGE_XMLNS,
				Ns3:  COVID_RESPONSE_XLMNS,
				Ns4:  DIGILOCKER_XLMNS,
				Req: &models2.Request{
					Text: "",
					ReqInfo: &models2.RequestInfo{
						MessageId:   uuid.New().String(),
						MessageDate: time.Now().Format("2006-01-02T15:04:05Z07:00"),
						ServiceId:   service.ServiceId,
						Sender: &models2.SenderCred{
							SenderId: a.Config.Shep.SenderLogin,
							Password: a.Config.Shep.SenderPassword,
						},
					},
					ReqData: &models2.RequestData{
						Text: "",
						Data: &models2.Data{
							Ns6:      COVID_REQUEST_XLMNS,
							Xsi:      XSI_XMLNS_SCEMA,
							Type:     COVID_REQUEST_TYPE,
							Iin:      docRequest.Iin,
							Login:    a.Config.Shep.ShepLogin,
							Password: a.Config.Shep.ShepPassword,
						},
					},
				},
			},
		},
	}

	shepResponse := &models2.EnvelopeResponse{}
	b, err := xml.Marshal(envelope)
	if err != nil {
		return shepResponse, err
	}
	req, err := http.NewRequest(http.MethodPost, a.Config.Shep.ShepEndpoint, bytes.NewBuffer(b))
	if err != nil {
		return shepResponse, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return shepResponse, err
	}

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("could not read response body")
		return shepResponse, err
	}

	err = xml.Unmarshal(response, &shepResponse)
	if err != nil {
		fmt.Println(err)
	}

	return shepResponse, nil
}
