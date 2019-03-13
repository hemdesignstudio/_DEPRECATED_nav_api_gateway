package request

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"github.com/hem-nav-gateway/config"
	"log"
	"net/http"
)

const (
	HTTP_METHOD = "POST"
)

type Request struct {
	Payload    []byte
	Uri        string
	SoapAction string
	Response   interface{}
}

func (r *Request) SoapRequest() {
	req, err := http.NewRequest(HTTP_METHOD, r.Uri, bytes.NewReader(r.Payload))
	if err != nil {
		log.Fatal("Error on creating request object. ", err.Error())
		return
	}
	req.Header.Set("Content-type", "text/xml")
	req.Header.Set("SOAPAction", r.SoapAction)
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
		log.Fatal("Error on dispatching request. ", err.Error())
		return
	}

	err = xml.NewDecoder(res.Body).Decode(r.Response)
	if err != nil {
		log.Fatal("Error on unmarshaling xml. ", err.Error())
		return
	}

	fmt.Println(res)
}
