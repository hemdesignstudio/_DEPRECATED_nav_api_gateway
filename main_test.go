package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootEndpoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	handler := RootEndpointHandler()
	handler.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "Response code is 200 as expected")
}
