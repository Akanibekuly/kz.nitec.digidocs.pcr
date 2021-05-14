package handler

import (
	"dd-pcr/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a *App) Process(c *gin.Context) {
	request := models.DocumentRequest{}
	c.BindJSON(&request)
	fmt.Printf("iin: %s\n", request.Iin)
	data, err := a.SendMessage(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.XML(http.StatusOK, data)
}

func (a *App) Ping(c *gin.Context) {
	c.String(200, "Pong")
}
