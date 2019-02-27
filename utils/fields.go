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

var types = map[string]*graphql.Object{
	"customer":     customer.CreateType(),
	"assemblyBom":  assemblybom.CreateType(),
	"item":         item.CreateType(),
	"salesOrder":   salesorder.CreateType(),
	"salesLine":    salesline.CreateType("SalesLine"),
	"postShip":     postship.CreateType(),
	"salesInvoice": salesinvoice.CreateType(),
}

func QueryType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
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
	})
}

func MutationType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"CreateCustomerCard": createFields("customer", CustomerCardArgs, customer.Create),
			"CreateItemCard":     createFields("item", itemCardArgs, item.Create),
			"CreateSalesOrder":   createFields("salesOrder", salesOrderArgs, salesorder.Create),
			"CreateSalesLine":    createFields("salesLine", createSalesLineArgs(), salesline.Create),
			"CreateSalesInvoice": createFields("salesInvoice", salesInvoiceArgs(), salesinvoice.Create),
			"UpdateCustomerCard": updateFields("customer", CustomerCardArgs, customer.Update),
			"UpdateItemCard":     updateFields("item", itemCardArgs, item.Update),
			"UpdateSalesOrder":   updateFields("salesOrder", salesOrderArgs, salesorder.Update),
			"UpdateSalesLine":    updateFields("salesLine", updateSalesLineArgs(), salesline.Update),
			"UpdateSalesInvoice": updateFields("salesInvoice", updateSalesInvoiceArgs(), salesinvoice.Update),
		},
	})
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

func createFields(
	name string,
	args map[string]*graphql.ArgumentConfig, create callbackWithArgs) *graphql.Field {

	field := &graphql.Field{
		Type: types[name],
		Args: args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return create(p.Args)
		},
	}
	return field
}

func updateFields(name string, args map[string]*graphql.ArgumentConfig, update callbackWithArgs) *graphql.Field {
	field := &graphql.Field{
		Type: types[name],
		Args: args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return update(p.Args)
		},
	}
	return field
}
