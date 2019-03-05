package salesinvoice

import (
	"github.com/graphql-go/graphql"
	"github.com/hem-nav-gateway/salesline"
)

var typeList = map[string]*graphql.Object{
	"salesLine": salesline.CreateType("Invoice_SalesLines"),
}

func getSalesLinesFields() *graphql.Field {
	field := &graphql.Field{
		Type: graphql.NewList(typeList["salesLine"]),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			salesInvoice := p.Source.(map[string]interface{})
			p.Args["key"] = "Document_No"
			p.Args["value"] = salesInvoice["No"]
			return salesline.Filter(nil, p.Args)
		},
	}
	return field
}

func extraFields() graphql.Fields {
	fields := graphql.Fields{
		"Sales_Lines": getSalesLinesFields(),
	}
	return fields
}
