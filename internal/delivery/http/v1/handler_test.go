// +build unit

package http

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPong(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	r.GET("/ping", Pong)

	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)

	assert := assert.New(t)
	assert.Equal(w.Code, http.StatusOK)

	p, err := ioutil.ReadAll(w.Body)
	if err != nil {
		assert.Fail("Response Body wrong")
	}

	assert.Equal(string(p), "pong")
}
