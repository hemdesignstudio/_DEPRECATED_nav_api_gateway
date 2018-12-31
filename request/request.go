package request

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"projects/graphql/config"
)

func GET(name string, url string) ([]byte, error) {
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
