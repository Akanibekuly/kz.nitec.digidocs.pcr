// +build unit

package Service

import (
	"encoding/xml"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	config3 "kz.nitec.digidocs.pcr/internal/config"
	models2 "kz.nitec.digidocs.pcr/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Mock() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		req, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		request := &models2.DocumentRequest{}
		err = xml.Unmarshal(req, request)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		s := ""
		i, err := w.Write([]byte(s))
		if err != nil || i != len(s) {
			return
		}
	})
	return r
}

func TestApp_SendMessage(t *testing.T) {
	assert := assert.New(t)
	config := config3.GetConfig()
	srv := httptest.NewServer(Mock())
	defer srv.Close()
	// заменяю на url тестового сервера
	config.Shep.ShepEndpoint = srv.URL
	app := App{
		Config: config,
	}
	serviceDTO := models2.ServiceDTO{
		Code:      "PcrCertificate",
		ServiceId: "CovidResult",
		Url:       "http://localhost:8095/pcr-cert",
	}
	documentDTO := models2.DocumentTypeDto{
		Code:   "PcrCertificate",
		NameEn: "The Result of PCR testing on COVID-19",
		NameKk: "COVID-19-ға тестілеу бойынша ПТР нәтижесі",
		NameRu: "Результат ПЦР тестирования на COVID-19",
	}
	cases := []struct {
		iin       string
		err       error
		expErr    error
		expResult *models2.EnvelopeResponse
	}{
		{
			iin:       "950110350170",
			err:       nil,
			expResult: &models2.EnvelopeResponse{},
		},
	}

	for _, v := range cases {
		request := models2.DocumentRequest{
			Iin: v.iin,
			Services: map[string]models2.ServiceDTO{
				"PCR_CERTIFICATE": serviceDTO,
			},
			DocumentTypeDto: documentDTO,
		}
		response, err := app.SendMessage(&request)
		assert.Equal(err, v.expErr)
		assert.Equal(response, v.expResult)
	}
}
