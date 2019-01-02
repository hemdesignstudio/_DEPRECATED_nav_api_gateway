package company

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"log"
	"projects/graphql/assemblybom"
	"projects/graphql/config"
	"projects/graphql/customer"
	"projects/graphql/item"
	"projects/graphql/request"
	"projects/graphql/salesorder"
)

type Company struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

func CreateCompanyType(

	customerType *graphql.Object,
	assemblybomType *graphql.Object,
	itemType *graphql.Object,
	salesOrderType *graphql.Object,

) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Company",
		Fields: graphql.Fields{
			"id":          &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"name":        &graphql.Field{Type: graphql.String},
			"displayName": &graphql.Field{Type: graphql.String},

			"AllCustomerCards": &graphql.Field{
				Type: graphql.NewList(customerType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					company, _ := p.Source.(*Company)
					log.Printf("fetching customers of company: %s", company.Name)
					return customer.GetCustomerCardByCompanyName(company.Name)
				},
			},

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

			"AllAssemblyBom": &graphql.Field{
				Type: graphql.NewList(assemblybomType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					company, _ := p.Source.(*Company)
					log.Printf("fetching Assembly BOM of company: %s", company.Name)
					return assemblybom.GetAssemblyBomByCompanyName(company.Name)
				},
			},

			"AssemblyBom": &graphql.Field{
				Type: assemblybomType,
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
				Type: graphql.NewList(itemType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					company, _ := p.Source.(*Company)
					log.Printf("fetching items of company: %s", company.Name)
					return item.GetItemCardByCompanyName(company.Name)
				},
			},

			"ItemCard": &graphql.Field{
				Type: itemType,
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
					return item.GetItemCardByNo(company.Name, no)
				},
			},

			"AllSalesOrders": &graphql.Field{
				Type: graphql.NewList(salesOrderType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					company, _ := p.Source.(*Company)
					log.Printf("fetching items of company: %s", company.Name)
					return salesorder.GetSalesOrderByCompanyName(company.Name)
				},
			},

			"SalesOrder": &graphql.Field{
				Type: salesOrderType,
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
					return salesorder.GetSalesOrderByNo(company.Name, no)
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
