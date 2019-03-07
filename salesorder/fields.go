// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package salesorder

import (
	"github.com/graphql-go/graphql"
	"github.com/hem-nav-gateway/salesline"
	"github.com/hem-nav-gateway/types"
)

// typeList Creates a Sales Line GraphQl Type
// which will be used as a nested field for SalesOrder
var typeList = map[string]*graphql.Object{
	"salesLine": salesline.CreateType("Order_SalesLines"),
}

// getSalesLinesFields creates a Salesline field objects
// which will be used as a nested field for SalesOrder
func getSalesLinesFields() *graphql.Field {
	field := &graphql.Field{
		Type: graphql.NewList(typeList["salesLine"]),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			salesOrder, ok := p.Source.(map[string]interface{})

			if ok == true {
				p.Args["key"] = "Document_No"
				p.Args["value"] = salesOrder["No"]
				obj := types.RequestObject{Args: p.Args}
				return salesline.Filter(obj)
			}
			return nil, nil

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
