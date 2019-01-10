package request

import (
	"errors"
	"fmt"
	"github.com/nav-api-gateway/config"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
)

func GET(uri string) ([]byte, error) {
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
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s: %s", "could not fetch data", resp.Status)
	}
	resultByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("could not read data")
	}

	return resultByte, nil
}

func GetAll(companyName string, endpoint string) []byte {
	uri := config.BaseUrl + config.CompanyEndpoint + fmt.Sprintf("('%s')", companyName) + endpoint
	resultByte, _ := GET(uri)

	return resultByte
}

func Filter(companyName, endpoint string, args map[string]interface{}) []byte {
	uri := config.BaseUrl + config.CompanyEndpoint + fmt.Sprintf("('%s')", companyName) + endpoint
	key := args["key"]
	value := args["value"]
	valueType := reflect.TypeOf(args["value"]).Kind()
	filter := fmt.Sprintf("?$filter=%s eq %s", key, value)

	if valueType == reflect.String {
		filter = fmt.Sprintf("?$filter=%s eq '%s'", key, value)
	}

	resultByte, _ := GET(uri + filter)
	return resultByte
}
