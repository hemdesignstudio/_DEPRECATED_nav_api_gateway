package company

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/nav-api-gateway/assemblybom"
	"github.com/nav-api-gateway/config"
	"github.com/nav-api-gateway/customer"
	"github.com/nav-api-gateway/item"
	"github.com/nav-api-gateway/request"
	"github.com/nav-api-gateway/salesline"
	"github.com/nav-api-gateway/salesorder"
	"log"
)

type Company struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

func CreateCompanyType(args map[string]*graphql.Object) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Company",
		Fields: graphql.Fields{
			"id":          &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"name":        &graphql.Field{Type: graphql.String},
			"displayName": &graphql.Field{Type: graphql.String},

			"AllCustomerCards": &graphql.Field{
				Type: graphql.NewList(args["customerType"]),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					company, _ := p.Source.(*Company)
					log.Printf("fetching customers of company: %s", company.Name)
					return customer.GetCustomerCardByCompanyName(company.Name)
				},
			},

			"CustomerCard": &graphql.Field{
				Type: args["customerType"],
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

			"AllAssemblyBom": &graphql.Field{
				Type: graphql.NewList(args["assemblyBomType"]),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					company, _ := p.Source.(*Company)
					log.Printf("fetching Assembly BOM of company: %s", company.Name)
					return assemblybom.GetAssemblyBomByCompanyName(company.Name)
				},
			},

			"AssemblyBom": &graphql.Field{
				Type: args["assemblyBomType"],
				Args: graphql.FieldConfigArgument{
					"no": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					company, _ := p.Source.(*Company)
					log.Printf("fetching AssemblyBom of company: %s", company.Name)
					name := p.Args["no"]
					no, _ := name.(string)
					return assemblybom.GetAssemblyByNo(company.Name, no)
				},
			},

			"AllItemCards": &graphql.Field{
				Type: graphql.NewList(args["itemType"]),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					company, _ := p.Source.(*Company)
					log.Printf("fetching items of company: %s", company.Name)
					return item.GetItemCardByCompanyName(company.Name)
				},
			},

			"ItemCard": &graphql.Field{
				Type: args["itemType"],
				Args: graphql.FieldConfigArgument{
					"no": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					company, _ := p.Source.(*Company)
					log.Printf("fetching item cards of company: %s", company.Name)
					name := p.Args["no"]
					no, _ := name.(string)
					return item.GetItemCardByNo(company.Name, no)
				},
			},

			"AllSalesOrders": &graphql.Field{
				Type: graphql.NewList(args["salesOrderType"]),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					company, _ := p.Source.(*Company)
					log.Printf("fetching sales orders of company: %s", company.Name)
					return salesorder.GetSalesOrderByCompanyName(company.Name)
				},
			},

			"AllSalesLines": &graphql.Field{
				Type: graphql.NewList(args["salesLineType"]),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					company, _ := p.Source.(*Company)
					log.Printf("fetching sales lines of company: %s", company.Name)
					return salesline.GetSalesLineByCompanyName(company.Name)
				},
			},
		},
	})
}

func GetCompanyByName(name string) (*Company, error) {
	conf := config.GetConfig()
	url := conf.BaseUrl + conf.CompanyEndpoint + fmt.Sprintf("('%s')", name)
	resultByte, err := request.GET(url)
	response := Company{}

	err = json.Unmarshal(resultByte, &response)
	if err != nil {
		return nil, errors.New("could not unmarshal data")
	}
	return &response, err
}
