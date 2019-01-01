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
	Id          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

func CreateCompanyType(customerType *graphql.Object) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Company",
		Fields: graphql.Fields{
			"id":          &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"name":        &graphql.Field{Type: graphql.String},
			"displayName": &graphql.Field{Type: graphql.String},

			"CustomerCard": &graphql.Field{
				Type: customerType,
				Args: graphql.FieldConfigArgument{
					"no": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					company, _ := p.Source.(*Company)
					log.Printf("fetching customers of company: %s", company.Name)
					name := p.Args["no"]
					no, _ := name.(string)
					return customer.GetCustomerCardByNo(company.Name, no)
				},
			},

			"AllCustomerCards": &graphql.Field{
				Type: graphql.NewList(customerType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					company, _ := p.Source.(*Company)
					log.Printf("fetching customers of company: %s", company.Name)
					return customer.GetCustomerCardByCompanyName(company.Name)
				},
			},
		},
	})
}

func GetCompanyByName(name string) (*Company, error) {
	conf := config.GetConfig()
	url := conf.BaseUrl + conf.CompanyEndpoint + fmt.Sprintf("('%s')", name)
	resultByte, err := request.GET(name, url)
	response := Company{}

	err = json.Unmarshal(resultByte, &response)
	if err != nil {
		return nil, errors.New("could not unmarshal data")
	}
	return &response, err
}
