package company

import (
	"github.com/graphql-go/graphql"
	"github.com/nav-api-gateway/assemblybom"
	"github.com/nav-api-gateway/config"
	"github.com/nav-api-gateway/customer"
	"github.com/nav-api-gateway/item"
	"github.com/nav-api-gateway/postship"
	"github.com/nav-api-gateway/salesinvoice"
	"github.com/nav-api-gateway/salesline"
	"github.com/nav-api-gateway/salesorder"
	"log"
)

var types = map[string]*graphql.Object{
	"customer":     customer.CreateType(),
	"assemblyBom":  assemblybom.CreateType(),
	"item":         item.CreateType(),
	"salesOrder":   salesorder.CreateType(),
	"salesLine":    salesline.CreateType("Company_SalesLine"),
	"postShip":     postship.CreateType(),
	"salesInvoice": salesinvoice.CreateType(),
}

func getCompanyFields() map[string]*graphql.Field {
	fields := map[string]*graphql.Field{
		"Id":          {Type: graphql.String},
		"Name":        {Type: graphql.String},
		"DisplayName": {Type: graphql.String},
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

func createCustomerCardFields() *graphql.Field {
	field := &graphql.Field{
		Type: types["customer"],
		Args: CustomerCardArgs,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			log.Printf("fetching Customer cards of company: %s", config.CompanyName)
			return customer.Create(p.Args)
		},
	}
	return field
}

func updateCustomerCardFields() *graphql.Field {
	field := &graphql.Field{
		Type: graphql.String,
		Args: CustomerCardArgs,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			log.Printf("fetching Customer cards of company: %s", config.CompanyName)
			return customer.Update(p.Args)
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

func createItemCardFields() *graphql.Field {
	field := &graphql.Field{
		Type: types["item"],
		Args: itemCardArgs,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			log.Printf("fetching Item cards of company: %s", config.CompanyName)
			return item.Create(p.Args)
		},
	}
	return field
}

func updateItemCardFields() *graphql.Field {
	field := &graphql.Field{
		Type: graphql.String,
		Args: itemCardArgs,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			log.Printf("update Item card of company: %s", config.CompanyName)
			return item.Update(p.Args)
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

func createSalesOrderFields() *graphql.Field {
	field := &graphql.Field{
		Type: types["salesOrder"],
		Args: salesOrderArgs,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			log.Printf("fetching Sales Order of company: %s", config.CompanyName)
			return salesorder.Create(p.Args)
		},
	}
	return field
}

func updateSalesOrderFields() *graphql.Field {
	field := &graphql.Field{
		Type: graphql.String,
		Args: salesOrderArgs,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			log.Printf("update Sales Order of company: %s", config.CompanyName)
			return salesorder.Update(p.Args)
		},
	}
	return field
}

func getSalesLinesFields() *graphql.Field {
	field := &graphql.Field{
		Type: graphql.NewList(types["salesLine"]),
		Args: filterArgs,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			log.Printf("fetching sales lines of company: %s", config.CompanyName)
			if len(p.Args) != 2 {
				return salesline.GetAll()
			}
			return salesline.Filter(p.Args)
		},
	}
	return field
}

func updateSalesLineFields() *graphql.Field {
	field := &graphql.Field{
		Type: graphql.String,
		Args: salesLineArgs,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			log.Printf("update Sales Lines of company: %s", config.CompanyName)
			return salesline.Update(p.Args)
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

func updatePostShipFields() *graphql.Field {
	field := &graphql.Field{
		Type: graphql.String,
		Args: postShipArgs,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			log.Printf("update Posted Sales Shipment of company: %s", config.CompanyName)
			return postship.Update(p.Args)
		},
	}
	return field
}

func getSalesInvoiceFields() *graphql.Field {
	field := &graphql.Field{
		Type: graphql.NewList(types["salesInvoice"]),
		Args: filterArgs,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			log.Printf("get sales invoices of company: %s", config.CompanyName)
			if len(p.Args) != 2 {
				return salesinvoice.GetAll()
			}
			return salesinvoice.Filter(p.Args)
		},
	}
	return field
}

func createSalesInvoiceFields() *graphql.Field {
	field := &graphql.Field{
		Type: types["salesInvoice"],
		Args: salesInvoiceArgs,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			log.Printf("fetching Sales Invoice of company: %s", config.CompanyName)
			return salesinvoice.Create(p.Args)
		},
	}
	return field
}

func updateSalesInvoiceFields() *graphql.Field {
	field := &graphql.Field{
		Type: graphql.String,
		Args: salesInvoiceArgs,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			log.Printf("update Sales Invoices of company: %s", config.CompanyName)
			return salesinvoice.Update(p.Args)
		},
	}
	return field
}
