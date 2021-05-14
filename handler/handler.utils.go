package handler

import (
	"bytes"
	"dd-pcr/models"
	"encoding/xml"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	ENVELOPE           = "Envelope"
	ENVELOP_SCHEMA     = "http://schemas.xmlsoap.org/soap/envelope/"
	SEND_MESSAGE_XMLNS = "http://bip.bee.kz/SyncChannel/v10/Types"
	COVID_RESPONSE_XLMNS = "http://api.nce.kz/SyncChannel/v1/Types/CovidResponse"
	DIGILOCKER_XLMNS = "http://digilocker.gov.kz/documentResponse/type/pcrcert"
	COVID_REQUEST_XLMNS = "http://api.nce.kz/SyncChannel/v1/Types/CovidRequest"
	XSI_XMLNS_SCEMA = "http://www.w3.org/2001/XMLSchema-instance"
	COVID_REQUEST_TYPE = "ns6:CovidRequest"
)

func (a *App) SendMessage(docRequest models.DocumentRequest) (models.EnvelopeResponse, error) {
	service:=docRequest.Services["PCR_CERTIFICATE"]
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
							SenderId: a.Config.Shep.SenderLogin,
							Password: a.Config.Shep.SenderPassword,
						},
					},
					ReqData: &models.RequestData{
						Text: "",
						Data: &models.Data{
							Ns6:      COVID_REQUEST_XLMNS,
							Xsi:      XSI_XMLNS_SCEMA,
							Type:     COVID_REQUEST_TYPE,
							Iin:      docRequest.Iin,
							Login:    "HDD1sQco26feoPeydupW",
							Password: "83297abe28d0a43b204c783af341430a",
						},
					},
				},
			},
		},
	}

	shepResponse:=&models.EnvelopeResponse{}
	b, err := xml.Marshal(envelope)
	if err != nil {
		return *shepResponse,err
	}
	req, err := http.NewRequest(http.MethodPost, a.Config.Shep.ShepEndpoint, bytes.NewBuffer(b))
	if err != nil {
		return *shepResponse,err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return *shepResponse,err
	}

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("could not read response body")
		return *shepResponse,err
	}

	err = xml.Unmarshal(response, &shepResponse)
	if err != nil {
		log.Println(err)
	}

	return *shepResponse, nil
}
