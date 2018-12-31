package customer

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"projects/graphql/config"
	"projects/graphql/request"
)

type Response struct {
	Value []Customer `json:"value"`
}

type Customer struct {
	No   string `json:"no"`
	Name string `json:"name"`
}

func CreateCustomerType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Customer",
		Fields: graphql.Fields{
			"no": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}

func FetchCustomersByCompanyName(name string) ([]Customer, error) {
	conf := config.GetConfig()
	url := conf.BaseUrl + conf.CompanyEndpoint + fmt.Sprintf("('%s')"+conf.CustomerCardWSEndpoint, name)
	resultByte, err := request.GET(name, url)
	res := Response{}
	err = json.Unmarshal(resultByte, &res)

	if err != nil {
		return nil, errors.New("could not unmarshal data")
	}
	return res.Value, nil
}
