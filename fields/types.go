// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

/*
Package fields implements a simple package for handling GraphQl fields.


*/
package fields

import (
	"github.com/graphql-go/graphql"
	"github.com/hem-nav-gateway/customer"
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
			//"AssemblyBom": queryFields(&assemblybom.Request{Company: company}),
			"CustomerCard": queryFields(&customer.Request{Company: company}),
			//"ItemCard":            queryFields(itemField(company)),
			//"SalesOrder":          queryFields(salesOrderField(company)),
			//"SalesLine":           queryFields(salesLineField(company)),
			//"PostedSalesShipment": queryFields(postShipField(company)),
			//"SalesInvoice":        queryFields(salesInvoiceField(company)),
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
			//"CreateItemCard":     createFields(itemField(company)),
			//"CreateSalesOrder":   createFields(salesOrderField(company)),
			//"CreateSalesLine":    createFields(salesLineField(company)),
			//"CreateSalesInvoice": createFields(salesInvoiceField(company)),

			//"UpdateCustomerCard": updateFields(customer.Request{Company: company}),
			//"UpdateItemCard":     updateFields(itemField(company)),
			//"UpdateSalesOrder":   updateFields(salesOrderField(company)),
			//"UpdateSalesLine":    updateFields(salesLineField(company)),
			//"UpdateSalesInvoice": updateFields(salesInvoiceField(company)),
		},
	}
	return graphql.NewObject(mutation)
}
