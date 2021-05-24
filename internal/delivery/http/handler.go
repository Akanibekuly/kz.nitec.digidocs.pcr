package http

import (
	"github.com/gin-gonic/gin"
	"kz.nitec.digidocs.pcr/internal/Service"
	"kz.nitec.digidocs.pcr/internal/models"
	"log"
	"net/http"
)

type Handler struct {
	Services *Service.Services
}

func NewHandler(services *Service.Services) *Handler {
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
		pcr.POST("/pcr-result", h.Process)
	}

	return router
}

func Pong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func (h *Handler) Process(c *gin.Context) {
	request := models.DocumentRequest{}
	err := c.BindJSON(&request)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal server error")
	}

	data, err := h.Services.PcrCertificateService.GetBySoap(request)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "Internal server error")
		return
	}

	var result *models.EnvelopeResponse
	switch data.(type) {
	case *models.EnvelopeResponse:
		result = data.(*models.EnvelopeResponse)
	default:
		log.Println("Unexpected data type")
		c.String(http.StatusInternalServerError, "Unexpected data type")
		return
	}

	c.JSON(http.StatusOK, result)

}
