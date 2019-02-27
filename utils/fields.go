package utils

import (
	"github.com/graphql-go/graphql"
	"github.com/nav-api-gateway/assemblybom"
	"github.com/nav-api-gateway/customer"
	"github.com/nav-api-gateway/item"
	"github.com/nav-api-gateway/postship"
	"github.com/nav-api-gateway/salesinvoice"
	"github.com/nav-api-gateway/salesline"
	"github.com/nav-api-gateway/salesorder"
)

type callback func() (interface{}, error)
type callbackWithArgs func(map[string]interface{}) (interface{}, error)
type Args map[string]map[string]*graphql.ArgumentConfig

var types = map[string]*graphql.Object{
	"assemblyBom":  assemblybom.CreateType(),
	"customer":     customer.CreateType(),
	"item":         item.CreateType(),
	"salesOrder":   salesorder.CreateType(),
	"salesLine":    salesline.CreateType("SalesLine"),
	"postShip":     postship.CreateType(),
	"salesInvoice": salesinvoice.CreateType(),
}

var args = Args{
	"customer":     customer.CreateArgs(),
	"item":         item.CreateArgs(),
	"salesOrder":   salesorder.CreateArgs(),
	"salesLine":    salesline.CreateArgs(),
	"postShip":     postship.CreateArgs(),
	"salesInvoice": salesinvoice.CreateArgs(),
}

func queryFields(name string, getAll callback, filter callbackWithArgs) *graphql.Field {
	field := &graphql.Field{
		Type: graphql.NewList(types[name]),
		Args: filterArgs,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if len(p.Args) != 2 {
				return getAll()
			}
			return filter(p.Args)
		},
	}
	return field
}

func createFields(name string, create callbackWithArgs) *graphql.Field {

	field := &graphql.Field{
		Type: types[name],
		Args: args[name],
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return create(p.Args)
		},
	}
	return field
}

func updateFields(name string, update callbackWithArgs) *graphql.Field {
	field := &graphql.Field{
		Type: types[name],
		Args: args[name],
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return update(p.Args)
		},
	}
	return field
}

func QueryType() *graphql.Object {
	query := graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"AssemblyBom":         queryFields("assemblyBom", assemblybom.GetAll, assemblybom.Filter),
			"CustomerCard":        queryFields("customer", customer.GetAll, customer.Filter),
			"ItemCard":            queryFields("item", item.GetAll, item.Filter),
			"SalesOrder":          queryFields("salesOrder", salesorder.GetAll, salesorder.Filter),
			"SalesLine":           queryFields("salesLine", salesline.GetAll, salesline.Filter),
			"PostedSalesShipment": queryFields("postShip", postship.GetAll, postship.Filter),
			"SalesInvoice":        queryFields("salesInvoice", salesinvoice.GetAll, salesinvoice.Filter),
		},
	}
	return graphql.NewObject(query)
}

func MutationType() *graphql.Object {
	mutation := graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"CreateCustomerCard": createFields("customer", customer.Create),
			"CreateItemCard":     createFields("item", item.Create),
			"CreateSalesOrder":   createFields("salesOrder", salesorder.Create),
			"CreateSalesLine":    createFields("salesLine", salesline.Create),
			"CreateSalesInvoice": createFields("salesInvoice", salesinvoice.Create),
			"UpdateCustomerCard": updateFields("customer", customer.Update),
			"UpdateItemCard":     updateFields("item", item.Update),
			"UpdateSalesOrder":   updateFields("salesOrder", salesorder.Update),
			"UpdateSalesLine":    updateFields("salesLine", salesline.Update),
			"UpdateSalesInvoice": updateFields("salesInvoice", salesinvoice.Update),
		},
	}

	return graphql.NewObject(mutation)
}
