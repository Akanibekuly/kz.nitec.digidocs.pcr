package http

import (
	"github.com/gin-gonic/gin"
	"kz.nitec.digidocs.pcr/internal/models"
	"kz.nitec.digidocs.pcr/internal/service"
	logs "kz.nitec.digidocs.pcr/pkg/logger"
	"kz.nitec.digidocs.pcr/pkg/utils"
	"log"
	"net/http"
)

type Handler struct {
	Services *service.Services
}

func NewHandler(services *service.Services) models.Handler {
	return &Handler{services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", Pong)

	pcr := router.Group("/digilocker/pcr-cert/api")
	{
		pcr.POST("/pcr-result", h.PcrTaskManager)
	}

	return router
}

func Pong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func (h *Handler) PcrTaskManager(c *gin.Context) {
	request := models.DocumentRequest{}

	err := c.BindJSON(&request)
	if err != nil {
		log.Printf("validation error %s", err)
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	logs.Logging(logs.GetRequestLog("INFO", "incoming request", "pcr_certificate", "dev_pcrtaskmananger", "", "", "", "", 12))
	//TODO request logging

	if !utils.CheckIin(request.Iin) {
		// TODO error logging
		log.Printf("Bad request: IIN %s doesn't correct\n", request.Iin)
		c.String(http.StatusBadRequest, "Bad request: iin %s", request.Iin)
		return
	}

	serviceInfo, err := h.Services.DocumentService.GetServiceInfoByCode()
	if err != nil {
		// TODO error logging
		log.Println(err)
		c.String(http.StatusInternalServerError, "Internal server error: %s", err)
		return
	}

	soapRequest := h.Services.PcrCertificateService.NewSoapRequest(&request, serviceInfo.ServiceId)

	data, err := h.Services.PcrCertificateService.GetBySoap(soapRequest, serviceInfo.URL)
	if err != nil {
		// TODO error logging
		log.Println(err)
		c.String(http.StatusInternalServerError, "Internal server error: %s", err)
		return
	}

	docInfo, err := h.Services.DocumentService.GetDocInfoByCode()
	if err != nil {
		// TODO error logging
		log.Println(err)
		c.String(http.StatusInternalServerError, "Internal server error: %s", err)
		return
	}

	//TODO docunent Request
	docResponse, err := h.Services.DocumentService.BuildDocumentResponse(docInfo, data)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal server error: %s", err)
		return
	}
	c.JSON(http.StatusOK, docResponse)
}
