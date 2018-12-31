package company

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"io/ioutil"
	"net/http"
	"projects/graphql/config"
)

var conf = config.GetConfig()

type Company struct {
	ID          string `json:id`
	Name        string `json:name`
	DisplayName string `json:displayName`
}

func CreateCompanyType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Company",
		Fields: graphql.Fields{

			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"displayName": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}

func GetCompanyByName(name string) (*Company, error) {
	var url = conf.BaseUrl + conf.CompanyEndpoint + fmt.Sprintf("('%s')", name)

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
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("could not read data")
	}
	result := Company{}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, errors.New("could not unmarshal data")
	}
	return &result, nil
}
