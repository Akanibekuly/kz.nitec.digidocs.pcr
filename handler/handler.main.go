package handler

import (
	"github.com/gin-gonic/gin"
)

func (a *App) Process(c *gin.Context){

}

func (a *App) Ping(c *gin.Context){
	c.String(200,"Pong")
}