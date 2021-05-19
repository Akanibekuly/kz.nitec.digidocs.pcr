package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	models2 "kz.nitec.digidocs.pcr/internal/models"
	"net/http"
)

func (a *App) Process(c *gin.Context) {
	request := models2.DocumentRequest{}
	c.BindJSON(&request)
	fmt.Printf("iin: %s\n", request.Iin)
	data, err := a.SendMessage(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.XML(http.StatusOK, *data)
}
