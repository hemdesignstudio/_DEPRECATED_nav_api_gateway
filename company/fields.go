package company

import (
	"github.com/graphql-go/graphql"
	"github.com/nav-api-gateway/assemblybom"
	"github.com/nav-api-gateway/customer"
	"github.com/nav-api-gateway/item"
	"github.com/nav-api-gateway/postship"
	"github.com/nav-api-gateway/salesline"
	"github.com/nav-api-gateway/salesorder"
	"log"
)

var types = map[string]*graphql.Object{
	"customer":    customer.CreateCustomerCardType(),
	"assemblyBom": assemblybom.CreateAssemblyBomType(),
	"item":        item.CreateItemCardType(),
	"salesOrder":  salesorder.CreateSalesOrderType(),
	"salesLine":   salesline.CreateSalesLineType(),
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
			company, _ := p.Source.(*Company)
			log.Printf("fetching Assembly BOM of company: %s", company.Name)
			if len(p.Args) != 2 {
				return assemblybom.GetAssemblyBomByCompanyName(company.Name)
			}
			return assemblybom.GetAssemblyBomByFilter(company.Name, p.Args)
		},
	}
	return field
}

func getCustomerCardFields() *graphql.Field {
	field := &graphql.Field{
		Type: graphql.NewList(types["customer"]),
		Args: filterArgs,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			company, _ := p.Source.(*Company)
			log.Printf("fetching Customer cards of company: %s", company.Name)
			if len(p.Args) != 2 {
				return customer.GetCustomerCardByCompanyName(company.Name)
			}
			return customer.GetCustomerCardByFilter(company.Name, p.Args)
		},
	}
	return field
}

func getItemCardFields() *graphql.Field {
	field := &graphql.Field{
		Type: graphql.NewList(types["item"]),
		Args: filterArgs,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			company, _ := p.Source.(*Company)
			log.Printf("fetching Items of company: %s", company.Name)
			if len(p.Args) != 2 {
				return item.GetItemCardByCompanyName(company.Name)
			}
			return item.GetItemCardByFilter(company.Name, p.Args)
		},
	}
	return field
}

func getSalesOrdersFields() *graphql.Field {
	field := &graphql.Field{
		Type: graphql.NewList(types["salesOrder"]),
		Args: filterArgs,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			company, _ := p.Source.(*Company)
			log.Printf("fetching sales orders of company: %s", company.Name)
			if len(p.Args) != 2 {
				return salesorder.GetSalesOrderByCompanyName(company.Name)
			}
			return salesorder.GetSalesOrderByFilter(company.Name, p.Args)
		},
	}
	return field
}

func getSalesLinesFields() *graphql.Field {
	field := &graphql.Field{
		Type: graphql.NewList(types["salesLine"]),
		Args: filterArgs,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			company, _ := p.Source.(*Company)
			log.Printf("fetching sales lines of company: %s", company.Name)
			if len(p.Args) != 2 {
				return salesline.GetSalesLineByCompanyName(company.Name)
			}

			return salesline.GetSalesLineByFilter(company.Name, p.Args)
		},
	}
	return field
}

func getPostShipFields() *graphql.Field {
	field := &graphql.Field{
		Type: graphql.NewList(types["postShip"]),
		Args: filterArgs,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			company, _ := p.Source.(*Company)
			log.Printf("fetching Posted sales shipments of company: %s", company.Name)
			if len(p.Args) != 2 {
				return postship.GetAll(company.Name)
			}
			return postship.Filter(company.Name, p.Args)
		},
	}
	return field
}
