// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

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

type callback func(interface{}) (interface{}, error)

type callbackWithArgs func(interface{}, interface{}) (interface{}, error)

/*
 types contains a map of GraphQL Type objects of (assemblybom, Customer, item  .. etc)

Example GraphQL type for 'AssemblyBom':

	Example:
		'''
		graphql.NewObject(graphql.ObjectConfig{
				Name: "AssemblyBom",
				Fields: graphql.Fields{
					"Parent_Item_No":       &graphql.Field{Type: graphql.String},
					"No":                   &graphql.Field{Type: graphql.String},
					"Type":                 &graphql.Field{Type: graphql.String},
					...
				},
			})
		'''


*/
var types = map[string]*graphql.Object{
	"assemblyBom":  assemblybom.CreateType(),
	"customer":     customer.CreateType(),
	"item":         item.CreateType(),
	"salesOrder":   salesorder.CreateType(),
	"salesLine":    salesline.CreateType("SalesLine"),
	"postShip":     postship.CreateType(),
	"salesInvoice": salesinvoice.CreateType(),
}

//filterArgs hold arguments used for filtering
var filterArgs = map[string]*graphql.ArgumentConfig{
	"key":   {Type: graphql.String},
	"value": {Type: graphql.String},
}

var companyArgs = map[string]*graphql.ArgumentConfig{
	"name": {Type: graphql.String},
}

/*
args hold all create and update arguments for all mutation types

Example of GraphQl Argument Object for 'customer'

	Example:

		customer.CreateArgs() would resolve to the following

		'''
		map[string]*graphql.ArgumentConfig{
			"No":			&graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"Name":			&graphql.ArgumentConfig{Type: graphql.String},
			"Address":		&graphql.ArgumentConfig{Type: graphql.String},
			...
		}
		'''

	Hint:
		arguments are used to create or update entities,
		some arguments are required and hence in the CustomerCard type,
		tags can be noticed

		example of required fields

			No	string `json:"No" required:"true"`

		and this will be translated to

			"No":	&graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},


The returned GraphQl arguments will be used as a part of the main mutation
*/
type Args map[string]map[string]*graphql.ArgumentConfig

var args = Args{
	"customer":     customer.CreateArgs(),
	"item":         item.CreateArgs(),
	"salesOrder":   salesorder.CreateArgs(),
	"salesLine":    salesline.CreateArgs(),
	"salesInvoice": salesinvoice.CreateArgs(),
}

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
func CompanyQueryType() *graphql.Object {
	query := graphql.ObjectConfig{
		Name: "Company",
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
func CompanyMutationType() *graphql.Object {
	mutation := graphql.ObjectConfig{
		Name: "CompanyMutation",
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

func QueryType() *graphql.Object {
	query := graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"Company": &graphql.Field{
				Type: CompanyQueryType(),
				Args: companyArgs,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					company := p.Args["name"]
					return company, nil
				},
			},
		},
	}
	return graphql.NewObject(query)
}

func MutationType() *graphql.Object {
	mutation := graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"Company": &graphql.Field{
				Type: CompanyMutationType(),
				Args: companyArgs,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					company := p.Args["name"]
					return company, nil
				},
			},
		},
	}
	return graphql.NewObject(mutation)
}
