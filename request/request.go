package request

import (
	"errors"
	"fmt"
	"github.com/nav-api-gateway/config"
	"io/ioutil"
	"net/http"
)

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
