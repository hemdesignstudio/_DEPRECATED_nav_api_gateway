package inventory

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"github.com/hem-nav-gateway/config"
	"io/ioutil"
	"net/http"
)

const (
	HTTP_METHOD  = "POST"
	SOAP_ACTION  = "urn:GetInventory"
	ENDPOINT_URI = "https://wsmyway182.navmoln.se:18201/MyWay182Services/WS/Hem%20TEST/Codeunit/GetInventory"
)

func soapRequest(payload []byte) (interface{}, error) {
	req, err := http.NewRequest(HTTP_METHOD, ENDPOINT_URI, bytes.NewReader(payload))
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

	//err = xml.NewDecoder(res.Body).Decode(&r.Response)
	byteValue, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err

	}

	//byteValue = exampleXML
	var response Response
	err = xml.Unmarshal(byteValue, &response)

	if err != nil {
		return nil, err

	}

	return &response.Body.GetInventoryResult.RetValues.InventoryLine, nil
}
