package utils

import (
	"bytes"
	"encoding/json"
	"github.com/hem-nav-gateway/roothandler"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func Client(method, query string, body interface{}) (int, []byte) {
	jsonBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, query, bytes.NewBuffer(jsonBody))
	response := httptest.NewRecorder()
	handler := roothandler.Handler("test")
	handler.ServeHTTP(response, request)
	resBody, _ := ioutil.ReadAll(response.Body)
	return response.Code, resBody
}
