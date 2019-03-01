package request

import (
	"bytes"
	"github.com/hem-nav-gateway/config"
	"github.com/hem-nav-gateway/errorhandler"
	"io/ioutil"
	"net/http"
	"net/url"
)

func headers(uri string, method string, body interface{}) *http.Request {
	u, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	u.RawQuery = u.Query().Encode()
	_body := []byte{}

	if body != nil {
		_body = body.([]byte)
	} else {
		_body = nil
	}

	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(_body))
	req.SetBasicAuth(config.Username, config.Passwd)
	req.Header.Add("If-Match", "*")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json; odata.metadata=minimal")
	return req

}

func clientRequest(req *http.Request, method string) (int, interface{}, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return resp.StatusCode, nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)

	if method == "POST" {
		if resp.StatusCode != http.StatusCreated {
			return resp.StatusCode, nil, errorhandler.Handle(resp.StatusCode, respBody)
		}
	} else if method == "DELETE" {
		if resp.StatusCode != http.StatusNoContent {
			return resp.StatusCode, nil, errorhandler.Handle(resp.StatusCode, respBody)
		}

	} else {
		if resp.StatusCode != http.StatusOK {
			return resp.StatusCode, nil, errorhandler.Handle(resp.StatusCode, respBody)

		}
	}
	return resp.StatusCode, respBody, nil
}

func request(uri string, method string, body interface{}) (int, interface{}, error) {
	req := headers(uri, method, body)
	return clientRequest(req, method)
}
