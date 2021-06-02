package models

import "github.com/gin-gonic/gin"

type Handler interface {
	TaskManager(c *gin.Context)
	InitRoutes() *gin.Engine
}
