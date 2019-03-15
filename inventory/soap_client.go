package inventory

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"github.com/hem-nav-gateway/config"
	"net/http"
	"net/url"
)

const (
	TYPE = "Codeunit"
)

func formatUri(company string) string {

	uri := fmt.Sprintf("%s/%s/%s%s", baseUrl, company, TYPE, endpoint)
	return uri
}

func soapRequest(obj *SoapObject) (interface{}, error) {

	uri := formatUri(obj.Company)
	u, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	u.RawQuery = u.Query().Encode()

	req, err := http.NewRequest(HTTP_METHOD, u.String(), bytes.NewReader(obj.Payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-type", "text/xml")
	req.Header.Set("SOAPAction", SOAP_ACTION)
	req.SetBasicAuth(config.Username, config.Passwd)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err

	}

	var response Response
	err = xml.NewDecoder(res.Body).Decode(&response)

	if err != nil {
		return nil, err

	}

	return &response.Body.GetInventoryResult.RetValues.InventoryLine, nil
}
