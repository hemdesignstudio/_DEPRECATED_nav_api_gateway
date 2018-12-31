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
	Value []CustomerCard `json:"value"`
}

type CustomerCard struct {
	No   string `json:"no"`
	Name string `json:"name"`
}

func CreateCustomerCardType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "CustomerCard",
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

func GetCustomerCardByCompanyName(name string) ([]CustomerCard, error) {
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

func GetCustomerCardByNo(companyName string, no string) (*CustomerCard, error) {
	conf := config.GetConfig()
	url := conf.BaseUrl +
		conf.CompanyEndpoint + fmt.Sprintf("('%s')", companyName) +
		conf.CustomerCardWSEndpoint + fmt.Sprintf("('%s')", no)

	resultByte, err := request.GET(no, url)
	res := CustomerCard{}
	err = json.Unmarshal(resultByte, &res)

	if err != nil {
		return nil, errors.New("could not unmarshal data")
	}
	return &res, nil
}
