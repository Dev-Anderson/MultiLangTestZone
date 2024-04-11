package main

import (
	"anderbank/controllers"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestRouterHome(t *testing.T) {
	mockResponse := `{"message":"Teste de home"}`
	r := SetUpRouter()
	r.GET("/api/v1/home", controllers.Home)
	req, _ := http.NewRequest("GET", "/api/v1/home", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
