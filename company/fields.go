package company

import (
	"github.com/graphql-go/graphql"
	"github.com/nav-api-gateway/assemblybom"
	"github.com/nav-api-gateway/config"
	"github.com/nav-api-gateway/customer"
	"github.com/nav-api-gateway/item"
	"github.com/nav-api-gateway/postship"
	"github.com/nav-api-gateway/salesorder"
	"log"
)

var types = map[string]*graphql.Object{
	"customer":    customer.CreateCustomerCardType(),
	"assemblyBom": assemblybom.CreateAssemblyBomType(),
	"item":        item.CreateItemCardType(),
	"salesOrder":  salesorder.CreateSalesOrderType(),
	"postShip":    postship.CreateType(),
}

var filterArgs = map[string]*graphql.ArgumentConfig{
	"key":   {Type: graphql.String},
	"value": {Type: graphql.String},
}

func getCompanyFields() map[string]*graphql.Field {
	fields := map[string]*graphql.Field{
		"id":          {Type: graphql.String},
		"name":        {Type: graphql.String},
		"displayName": {Type: graphql.String},
	}
	return fields
}

func getAssemblyBomFields() *graphql.Field {
	field := &graphql.Field{
		Type: graphql.NewList(types["assemblyBom"]),
		Args: filterArgs,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			log.Printf("fetching Assembly BOM of company: %s", config.CompanyName)
			if len(p.Args) != 2 {
				return assemblybom.GetAll()
			}
			return assemblybom.Filter(p.Args)
		},
	}
	return field
}

func getCustomerCardFields() *graphql.Field {
	field := &graphql.Field{
		Type: graphql.NewList(types["customer"]),
		Args: filterArgs,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			log.Printf("fetching Customer cards of company: %s", config.CompanyName)
			if len(p.Args) != 2 {
				return customer.GetAll()
			}
			return customer.Filter(p.Args)
		},
	}
	return field
}

func getItemCardFields() *graphql.Field {
	field := &graphql.Field{
		Type: graphql.NewList(types["item"]),
		Args: filterArgs,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			log.Printf("fetching Items of company: %s", config.CompanyName)
			if len(p.Args) != 2 {
				return item.GetAll()
			}
			return item.Filter(p.Args)
		},
	}
	return field
}

func getSalesOrdersFields() *graphql.Field {
	field := &graphql.Field{
		Type: graphql.NewList(types["salesOrder"]),
		Args: filterArgs,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			log.Printf("fetching sales orders of company: %s", config.CompanyName)
			if len(p.Args) != 2 {
				return salesorder.GetAll()
			}
			return salesorder.Filter(p.Args)
		},
	}
	return field
}

func getPostShipFields() *graphql.Field {
	field := &graphql.Field{
		Type: graphql.NewList(types["postShip"]),
		Args: filterArgs,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			log.Printf("fetching Posted sales shipments of company: %s", config.CompanyName)
			if len(p.Args) != 2 {
				return postship.GetAll()
			}
			return postship.Filter(p.Args)
		},
	}
	return field
}
