package request

import (
	"bytes"
	"errors"
	"github.com/nav-api-gateway/config"
	"github.com/nav-api-gateway/errorhandler"
	"io/ioutil"
	"net/http"
	"net/url"
)

func get(uri string) ([]byte, error) {
	u, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	u.RawQuery = u.Query().Encode()

	client := &http.Client{}
	req, err := http.NewRequest("GET", u.String(), nil)
	req.SetBasicAuth(config.Username, config.Passwd)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	resByte, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, errorhandler.Handle(resp.StatusCode, resByte)
	}
	return resByte, nil
}

func post(uri string, body []byte) ([]byte, error) {
	u, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	u.RawQuery = u.Query().Encode()

	client := &http.Client{}
	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(body))
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
	if resp.StatusCode != http.StatusCreated {
		return nil, errorhandler.Handle(resp.StatusCode, resultByte)
	}
	if err != nil {
		return nil, errors.New("could not read data")
	}
	return resultByte, nil
}

func patch(uri string, body []byte) ([]byte, error) {
	u, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	u.RawQuery = u.Query().Encode()

	client := &http.Client{}
	req, err := http.NewRequest("PATCH", u.String(), bytes.NewBuffer(body))
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
	if resp.StatusCode != http.StatusOK {
		return nil, errorhandler.Handle(resp.StatusCode, resultByte)
	}
	if err != nil {
		return nil, errors.New("could not read data")
	}
	return resultByte, nil
}
