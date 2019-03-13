/*
Package fields implements a simple package for handling GraphQl fields.


*/
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

/*
QueryType creates the root query with all of its nested fields

	Example:

		fields:
			"AssemblyBom",
			"CustomerCard",
			"ItemCard",
			"SalesOrder",
			"SalesLine",
			"PostedSalesShipment",
			"SalesInvoice",

		queryFields("assemblyBom", assemblybom.GetAll, assemblybom.Filter) would resolve to

			'''
			&graphql.Field{
				Type: graphql.NewList(types["assemblyBom"]),
				Args: filterArgs,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if len(p.Args) != 2 {
						return assemblybom.GetAll()
					}
					return assemblybom.Filter(p.Args)
				},
			}
			'''
*/
func QueryType(company string) *graphql.Object {
	query := graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"AssemblyBom":         queryFields(&assemblybom.Request{Company: company}),
			"CustomerCard":        queryFields(&customer.Request{Company: company}),
			"ItemCard":            queryFields(&item.Request{Company: company}),
			"SalesOrder":          queryFields(&salesorder.Request{Company: company}),
			"SalesLine":           queryFields(&salesline.Request{Company: company}),
			"PostedSalesShipment": queryFields(&postship.Request{Company: company}),
			"SalesInvoice":        queryFields(&salesinvoice.Request{Company: company}),
		},
	}
	return graphql.NewObject(query)
}

/*
MutationType create the root mutation (Create or updating an entity) for all types

	Example:

		fields:
			"CreateCustomerCard",
			"CreateItemCard",
			"CreateSalesOrder",
			"CreateSalesLine",
			"CreateSalesInvoice",
			"UpdateCustomerCard",
			"UpdateItemCard",
			"UpdateSalesOrder",
			"UpdateSalesLine",
			"UpdateSalesInvoice"

		createFields("customer", customer.Create) would resolve to

			'''
			&graphql.Field{
				Type: types["customer"],
				Args: CustomerCardArgs,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					log.Printf("fetching Customer cards of company: %s", config.CompanyName)
					return customer.Create(p.Args)
				},
			}
			'''

		updateFields("customer", customer.Update) would resolve to

			'''
			&graphql.Field{
				Type: types["customer"],
				Args: CustomerCardArgs,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					log.Printf("updating Customer cards of company: %s", config.CompanyName)
					return customer.Update(p.Args)
				},
			}
			'''
*/
func MutationType(company string) *graphql.Object {
	mutation := graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"CreateCustomerCard": createFields(&customer.Request{Company: company}),
			"CreateItemCard":     createFields(&item.Request{Company: company}),
			"CreateSalesOrder":   createFields(&salesorder.Request{Company: company}),
			"CreateSalesLine":    createFields(&salesline.Request{Company: company}),
			"CreateSalesInvoice": createFields(&salesinvoice.Request{Company: company}),
			"UpdateCustomerCard": updateFields(&customer.Request{Company: company}),
			"UpdateItemCard":     updateFields(&item.Request{Company: company}),
			"UpdateSalesOrder":   updateFields(&salesorder.Request{Company: company}),
			"UpdateSalesLine":    updateFields(&salesline.Request{Company: company}),
			"UpdateSalesInvoice": updateFields(&salesinvoice.Request{Company: company}),
		},
	}
	return graphql.NewObject(mutation)
}
