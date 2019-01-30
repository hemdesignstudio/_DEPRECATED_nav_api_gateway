package test

import (
	"github.com/nav-api-gateway/roothandler"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAssemblyBOM(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	handler := roothandler.RootEndpoint()
	handler.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "Response code is 200 as expected")
}
