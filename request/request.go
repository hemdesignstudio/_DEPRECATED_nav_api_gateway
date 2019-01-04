package request

import (
	"errors"
	"fmt"
	"github.com/nav-api-gateway/config"
	"io/ioutil"
	"net/http"
)

var conf = config.GetConfig()

func GET(url string) ([]byte, error) {
	conf := config.GetConfig()
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(conf.Username, conf.Passwd)
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
	url := conf.BaseUrl + conf.CompanyEndpoint + fmt.Sprintf("('%s')", companyName) + endpoint
	resultByte, _ := GET(url)

	return resultByte
}

func Filter(companyName, endpoint string, args map[string]interface{}) []byte {
	url := conf.BaseUrl + conf.CompanyEndpoint + fmt.Sprintf("('%s')", companyName) + endpoint
	filter := fmt.Sprintf("?$filter=%s+eq+'%s'", args["name"], args["value"])
	resultByte, _ := GET(url + filter)

	return resultByte
}
