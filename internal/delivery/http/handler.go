package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kz.nitec.digidocs.pcr/internal/Service"
	"kz.nitec.digidocs.pcr/internal/models"
	"net/http"
)

type Handler struct {
	Services *Service.Services
}

func NewHandler(services *Service.Services) *Handler{
	return  &Handler{services}
}

func (h *Handler) InitRoutes() *gin.Engine{
	router:=gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		)
	router.GET("/ping", Pong)
	pcr := router.Group("/digilocker/pcr-cert/api")
	{
		pcr.POST("/pcr-result")
	}

	return router
}

func Pong(c *gin.Context){
	c.String(http.StatusOK, "pong")
}


func (a *App) Process(c *gin.Context) {
	request := models.DocumentRequest{}
	c.BindJSON(&request)
	fmt.Printf("iin: %s\n", request.Iin)
	data, err := a.SendMessage(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.XML(http.StatusOK, *data)
}
