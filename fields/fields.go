package fields

import (
	"github.com/graphql-go/graphql"
	"github.com/hem-nav-gateway/assemblybom"
	"github.com/hem-nav-gateway/customer"
	"github.com/hem-nav-gateway/item"
	"github.com/hem-nav-gateway/postship"
	"github.com/hem-nav-gateway/salesinvoice"
	"github.com/hem-nav-gateway/salesline"
	"github.com/hem-nav-gateway/salesorder"
)

type callback func(interface{}) (interface{}, error)
type callbackWithArgs func(interface{}, interface{}) (interface{}, error)
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

var filterArgs = map[string]*graphql.ArgumentConfig{
	"key":   {Type: graphql.String},
	"value": {Type: graphql.String},
}

var args = Args{
	"customer":     customer.CreateArgs(),
	"item":         item.CreateArgs(),
	"salesOrder":   salesorder.CreateArgs(),
	"salesLine":    salesline.CreateArgs(),
	"salesInvoice": salesinvoice.CreateArgs(),
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
