package company

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"log"
	"projects/graphql/config"
	"projects/graphql/customer"
	"projects/graphql/request"
)

type Company struct {
	ID          string `json:id`
	Name        string `json:name`
	DisplayName string `json:displayName`
}

func CreateCompanyType(customerType *graphql.Object) *graphql.Object {
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
			"CustomerCardWS": &graphql.Field{
				Type: graphql.NewList(customerType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					company, _ := p.Source.(*Company)
					log.Printf("fetching customers of company: %s", company.Name)
					return customer.FetchCustomersByCompanyName(company.Name)
				},
			},
		},
	})
}

func GetCompanyByName(name string) (*Company, error) {
	conf := config.GetConfig()
	url := conf.BaseUrl + conf.CompanyEndpoint + fmt.Sprintf("('%s')", name)
	resultByte, err := request.GET(name, url)
	result := Company{}

	err = json.Unmarshal(resultByte, &result)
	if err != nil {
		return nil, errors.New("could not unmarshal data")
	}
	return &result, err
}
