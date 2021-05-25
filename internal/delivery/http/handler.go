package http

import (
	"github.com/gin-gonic/gin"
	"kz.nitec.digidocs.pcr/internal/models"
	"kz.nitec.digidocs.pcr/internal/service"
	logs "kz.nitec.digidocs.pcr/pkg/logger"
	"log"
	"net/http"
)

type Handler struct {
	Services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
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
		logs.Logging(logs.GetStandardLog("ERROR", "Bad request", "pcr-app", "INFO", ""))
		c.String(http.StatusInternalServerError, "Internal server error")
		return
	}

	serviceInfo, err := h.Services.DocumentService.GetServiceInfoByCode()
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal server error")
		return
	}

	soapRequest := h.Services.PcrCertificateService.NewSoapRequest(serviceInfo.ServiceId, &request)

	data, err := h.Services.PcrCertificateService.GetBySoap(soapRequest, serviceInfo.URL)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "Internal server error")
		return
	}

	c.JSON(http.StatusOK, data)
}
