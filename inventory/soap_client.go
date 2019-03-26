package inventory

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/xml"
	"fmt"
	"github.com/hem-nav-gateway/config"
	"io/ioutil"
	"log"
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

func handleTls() *http.Transport {

	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	certs, err := ioutil.ReadFile(config.CA_CERT)
	if err != nil {
		log.Fatalf("Failed to append %q to RootCAs: %v", config.CA_CERT, err)
	}

	if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
		log.Println("No certs appended, using system certs only")
	}

	conf := &tls.Config{
		InsecureSkipVerify: false,
		RootCAs:            rootCAs,
	}

	tr := &http.Transport{TLSClientConfig: conf}

	return tr
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
	username, password := config.NAVCredentials()
	req.Header.Set("Content-type", "text/xml")
	req.Header.Set("SOAPAction", SOAP_ACTION)
	req.SetBasicAuth(username, password)

	tr := handleTls()
	client := &http.Client{Transport: tr}

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
