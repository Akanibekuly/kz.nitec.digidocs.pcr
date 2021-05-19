package handler

import (
	"encoding/xml"
	"fmt"
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
		w.Write([]byte(""))
	})
	return r
}

func TestApp_SendMessage(t *testing.T) {
	config := config3.GetConfig()
	srv := httptest.NewServer(Mock())
	defer srv.Close()
	// заменяю на url тестового сервера
	config.Shep.ShepEndpoint = srv.URL
	app := App{
		Config: config,
	}
	serviceDTO := models2.ServiceDTO{
		Code:      "PCR_CERTIFICATE",
		ServiceId: "CovidResult",
		Url:       "http://localhost:8095/pcr-cert",
	}
	documentDTO := models2.DocumentTypeDto{
		Code:   "",
		NameEn: "",
		NameKk: "",
		NameRu: "",
	}
	cases := []struct {
		iin           string
		err           error
		expectedError error
	}{
		{iin: "950110350170"},
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
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(response)
	}
}
