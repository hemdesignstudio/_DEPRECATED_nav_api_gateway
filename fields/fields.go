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

type fieldType interface {
	GetName() string
	GetCompany() string
	SetArgs(map[string]interface{})
	SetFields([]string)
	GetAll() (interface{}, error)
	Filter() (interface{}, error)
	Update() (interface{}, error)
	Create() (interface{}, error)
}

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

type Fields map[string]*graphql.Field

var args = Args{
	"customer":     customer.CreateArgs(),
	"item":         item.CreateArgs(),
	"salesOrder":   salesorder.CreateArgs(),
	"salesLine":    salesline.CreateArgs(),
	"salesInvoice": salesinvoice.CreateArgs(),
}

//func assemblyBomField(company string) fieldType {
//	field := fieldType{
//		Name:    "assemblyBom",
//		Company: company,
//		Filter:  assemblybom.Filter,
//		GetAll:  assemblybom.GetAll,
//	}
//	return field
//}
//
//func customerField(company string) fieldType {
//	field := fieldType{
//		Name:    "customer",
//		Company: company,
//		Filter:  customer.Filter,
//		GetAll:  customer.GetAll,
//		Create:  customer.Create,
//		Update:  customer.Update,
//	}
//	return field
//}
//
//func itemField(company string) fieldType {
//	field := fieldType{
//		Name:    "item",
//		Company: company,
//		Filter:  item.Filter,
//		GetAll:  item.GetAll,
//		Create:  item.Create,
//		Update:  item.Update,
//	}
//	return field
//}
//
//func salesOrderField(company string) fieldType {
//	field := fieldType{
//		Name:    "salesOrder",
//		Company: company,
//		Filter:  salesorder.Filter,
//		GetAll:  salesorder.GetAll,
//		Create:  salesorder.Create,
//		Update:  salesorder.Update,
//	}
//	return field
//}
//
//func salesLineField(company string) fieldType {
//	field := fieldType{
//		Name:    "salesLine",
//		Company: company,
//		Filter:  salesline.Filter,
//		GetAll:  salesline.GetAll,
//		Create:  salesline.Create,
//		Update:  salesline.Update,
//	}
//	return field
//}
//
//func postShipField(company string) fieldType {
//	field := fieldType{
//		Name:    "postShip",
//		Company: company,
//		Filter:  postship.Filter,
//		GetAll:  postship.GetAll,
//	}
//	return field
//}
//
//func salesInvoiceField(company string) fieldType {
//	field := fieldType{
//		Name:    "salesInvoice",
//		Company: company,
//		Filter:  salesinvoice.Filter,
//		GetAll:  salesinvoice.GetAll,
//		Create:  salesinvoice.Create,
//		Update:  salesinvoice.Update,
//	}
//	return field
//}
