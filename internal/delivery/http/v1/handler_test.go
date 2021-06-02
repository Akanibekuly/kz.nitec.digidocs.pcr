// +build unit

package http

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"kz.nitec.digidocs.pcr/internal/models"
	"kz.nitec.digidocs.pcr/internal/service"
	"net/http"
	"net/http/httptest"
	"testing"
	"errors"
)

func TestPong(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	r.GET("/ping", Pong)

	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)

	assert := assert.New(t)
	assert.Equal(w.Code, http.StatusOK)

	p, err := ioutil.ReadAll(w.Body)
	if err != nil {
		assert.Fail("Response Body wrong")
	}

	assert.Equal(string(p), "pong")
}

func TestHandler_TaskManager(t *testing.T) {
	services := &service.Services{
		&MockPersonPhotoService{},
	}
	h := NewHandler(services)

	assert.New(t)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body := bytes.NewBuffer([]byte("{\"iin\":\"950110350172\",\n\"services\": {\n\"PCR_CERTIFICATE\":\n{\"code\":\"PCR_CERTIFICATE\",\n\"serviceId\":\"CovidResult\",\n\"url\": \"http://localhost:8095/pcr-cert\"}\n},\n\"documentType\": {\n\"code\": \"\",\n\"nameRu\":  \"nameRu\",\n\"nameKk\": \"nameKk\"}}\n"))
	c.Request, _ = http.NewRequest("POST", "/", body)
	h.TaskManager(c)

	assert := assert.New(t)
	assert.Equal(w.Code, 400)
}

type MockPersonPhotoService struct {
	mock.Mock
}

func (m *MockPersonPhotoService) GetBySoap(request *models.SoapRequest) (*models.SoapResponse, error) {
	args := m.Called(request)
	var response *models.SoapResponse
	switch args.Get(1).(type) {
	case *models.SoapResponse:
		response = args.Get(1).(*models.SoapResponse)
	default:
		return nil, errors.New("invalid type format")
	}
	return response, args.Error(1)
}
func (m *MockPersonPhotoService) NewSoapRequest(request *models.DocumentRequest) (*models.SoapRequest, error) {
	args := m.Called(request)
	var response *models.SoapRequest
	switch args.Get(1).(type) {
	case *models.SoapRequest:
		response = args.Get(1).(*models.SoapRequest)
	default:
		return nil, errors.New("invalid type format")
	}
	return response, args.Error(1)
}
