package test

import (
	"github.com/nav-api-gateway/roothandler"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func client(method, query string) (int, []byte) {
	request, _ := http.NewRequest(method, query, nil)
	response := httptest.NewRecorder()
	handler := roothandler.RootEndpoint()
	handler.ServeHTTP(response, request)
	resBody, _ := ioutil.ReadAll(response.Body)
	return response.Code, resBody

}
