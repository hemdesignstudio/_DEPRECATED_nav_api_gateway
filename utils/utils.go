package utils

import (
	"github.com/graphql-go/graphql"
)

func QueryType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"AssemblyBom":         getAssemblyBomFields(),
			"CustomerCard":        getCustomerCardFields(),
			"ItemCard":            getItemCardFields(),
			"SalesOrder":          getSalesOrdersFields(),
			"SalesLine":           getSalesLineFields(),
			"PostedSalesShipment": getPostShipFields(),
			"SalesInvoice":        getSalesInvoiceFields(),
		},
	})
}

func MutationType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"CreateCustomerCard": createCustomerCardFields(),
			"CreateItemCard":     createItemCardFields(),
			"CreateSalesOrder":   createSalesOrderFields(),
			"CreateSalesLine":    createSalesLineFields(),
			//"CreatePostedSalesShipment": createPostShipFields(),
			"CreateSalesInvoice": createSalesInvoiceFields(),
			"UpdateCustomerCard": updateCustomerCardFields(),
			"UpdateItemCard":     updateItemCardFields(),
			"UpdateSalesOrder":   updateSalesOrderFields(),
			"UpdateSalesLine":    updateSalesLineFields(),
			//"UpdatePostedSalesShipment": updatePostShipFields(),
			"UpdateSalesInvoice": updateSalesInvoiceFields(),
		},
	})
}
