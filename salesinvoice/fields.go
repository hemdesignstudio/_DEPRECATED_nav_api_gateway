package salesinvoice

import (
	"github.com/graphql-go/graphql"
	"github.com/nav-api-gateway/config"
	"github.com/nav-api-gateway/salesline"
	"log"
)

var types = map[string]*graphql.Object{
	"salesLine": salesline.CreateType("Invoice_SalesLines"),
}

func getSalesLinesFields() *graphql.Field {
	field := &graphql.Field{
		Type: graphql.NewList(types["salesLine"]),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			log.Printf("fetching sales lines of company: %s which are related to Sales Invoices", config.CompanyName)
			salesInvoice, _ := p.Source.(SalesInvoice)
			p.Args["key"] = "Document_No"
			p.Args["value"] = salesInvoice.No
			return salesline.Filter(p.Args)
		},
	}
	return field
}
