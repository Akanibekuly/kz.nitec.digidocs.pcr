package models

import "github.com/gin-gonic/gin"

type Handler interface {
	PcrTaskManager(c *gin.Context)
	InitRoutes() *gin.Engine
}
