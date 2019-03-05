// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package salesinvoice

import (
	"github.com/graphql-go/graphql"
	"github.com/hem-nav-gateway/salesline"
)

// typeList Creates a Sales Line GraphQl Type which will
// be used as a nested field for Sales Invoice
var typeList = map[string]*graphql.Object{
	"salesLine": salesline.CreateType("Invoice_SalesLines"),
}

// getSalesLinesFields creates a Salesline field objects
// which will be used as a nested field for Sales Invoice
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

// extraFields are nested fields from
// other Types (SalesLine, Item, Customer .. etc)
func extraFields() graphql.Fields {
	fields := graphql.Fields{
		"Sales_Lines": getSalesLinesFields(),
	}
	return fields
}
