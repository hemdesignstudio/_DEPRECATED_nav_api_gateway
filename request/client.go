package request

import (
	"bytes"
	"github.com/nav-api-gateway/config"
	"github.com/nav-api-gateway/errorhandler"
	"io/ioutil"
	"net/http"
	"net/url"
)

func clientRequest(uri string, method string, body []byte) ([]byte, error) {
	u, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	u.RawQuery = u.Query().Encode()
	client := &http.Client{}
	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(body))
	req.SetBasicAuth(config.Username, config.Passwd)
	req.Header.Add("If-Match", "*")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json; odata.metadata=minimal")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	resultByte, err := ioutil.ReadAll(resp.Body)

	if method == "POST" {
		if resp.StatusCode != http.StatusCreated {
			return nil, errorhandler.Handle(resp.StatusCode, resultByte)
		}
	} else {
		if resp.StatusCode != http.StatusOK {
			return nil, errorhandler.Handle(resp.StatusCode, resultByte)

		}
	}
	return resultByte, nil
}
