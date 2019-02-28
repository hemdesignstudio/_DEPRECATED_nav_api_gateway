package utils

import (
	"bytes"
	"encoding/json"
	"github.com/hem-nav-gateway/config"
	"github.com/hem-nav-gateway/roothandler"
	"github.com/hem-nav-gateway/session"

	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func Client(method, query string, body interface{}) (int, []byte) {
	jsonBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, query, bytes.NewBuffer(jsonBody))
	response := httptest.NewRecorder()
	session.SetSession(request, config.TestCompanyName)
	handler := roothandler.Handler()
	handler.ServeHTTP(response, request)
	resBody, _ := ioutil.ReadAll(response.Body)
	return response.Code, resBody
}
