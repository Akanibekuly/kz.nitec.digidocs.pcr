package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kz.nitec.digidocs.pcr/internal/models"
	"kz.nitec.digidocs.pcr/internal/service"
	"kz.nitec.digidocs.pcr/pkg/logger"
	"kz.nitec.digidocs.pcr/pkg/utils"
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
		pcr.POST("/pcr-result", h.TaskManager)
	}

	return router
}

func Pong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func (h *Handler) TaskManager(c *gin.Context) {
	request := models.DocumentRequest{}

	err := c.BindJSON(&request)
	if err != nil {
		logger.PrintLog("ERROR", "PCR_TM", "", logger.CreateMessageLog(err))
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	logger.PrintLog("INFO", "PCR_TM", "", fmt.Sprintf("Request from %s with iin: %s", "RabbitMq", request.Iin))

	if !utils.CheckIin(request.Iin) {
		logger.PrintLog("ERROR", "PCR_TM", "", fmt.Sprintf("Bad request: IIN %s doesn't correct\n", request.Iin))
		c.String(http.StatusBadRequest, "Bad request: iin %s", request.Iin)
		return
	}

	soapRequest, err := h.Services.PcrCertificateService.NewSoapRequest(&request)
	if err != nil {
		logger.PrintLog("ERROR", "PCR-TM", "", err)
		c.String(http.StatusInternalServerError, "Internal server error: %s", err)
		return
	}

	soapResponse, err := h.Services.PcrCertificateService.GetBySoap(soapRequest)
	if err != nil {
		logger.PrintLog("ERROR", "PCR-TM", "", err)
		c.String(http.StatusInternalServerError, "Internal server error: %s", err)
		return
	}

	result, err := h.Services.BuildService.BuildDocumentResponse(soapResponse)
	if err != nil {
		logger.PrintLog("ERROR", "PCR-TM", "", err)
		c.String(http.StatusInternalServerError, "Internal server error: %s", err)
		return
	}
	c.JSON(http.StatusOK, result)
}
