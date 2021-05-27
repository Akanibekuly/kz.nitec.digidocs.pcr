// +build unit

package http

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"kz.nitec.digidocs.pcr/internal/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

//func TestHandler_InitRoutes(t *testing.T) {
//	assert:=assert.New(t)
//	err:=os.Setenv("APP_MODE","debug")
//	if err!=nil{
//		assert.Fail("Cannot set env mode")
//	}
//	gin.SetMode(os.Getenv("APP_MODE"))
//	r:=gin.Default()
//	w:=httptest.NewRecorder()
//
//	h:=&Handler{
//		Services: &service.Services{
//
//		},
//	}
//
//
//
//
//	router:=h.InitRoutes()
//
//	err=os.Unsetenv("APP_MODE")
//	if err!=nil{
//		assert.Fail("Cannot unset env APP_MODE")
//	}
//}

func TestPong(t *testing.T) {
	w:=httptest.NewRecorder()
	r:=gin.Default()
	r.GET("/ping",Pong)

	req,_:=http.NewRequest("GET","/ping",nil)
	r.ServeHTTP(w,req)

	assert:=assert.New(t)
	assert.Equal(w.Code, http.StatusOK)

	p,err:=ioutil.ReadAll(w.Body)
	if err!=nil{
		assert.Fail("Response Body wrong")
	}

	assert.Equal(string(p),"pong")
}

func TestHandler_PcrTaskManager(t *testing.T) {
	services:=&service.Services{}
	h:=NewHandler(services)

	assert.New(t)
	r:=gin.Default()
	w:=httptest.NewRecorder()
	w.Write()
	c,err:=gin.CreateTestContext(w)
	if err!=nil{
		assert.Fail("Gin context error")
	}
	r.POST("/",h.PcrTaskManager())
	req,_:=httptest.NewRecorder()

}
