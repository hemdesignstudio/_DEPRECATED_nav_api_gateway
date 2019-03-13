// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package salesorder

import (
	"github.com/graphql-go/graphql"
	"github.com/hem-nav-gateway/salesline"
)

// typeList Creates a Sales Line GraphQl Type
// which will be used as a nested field for SalesOrder
var _salesLine = salesline.Request{}
var salesLineType = *_salesLine.NewType("Order_SalesLines")

// getSalesLinesFields creates a Salesline field objects
// which will be used as a nested field for SalesOrder
func getSalesLinesFields() *graphql.Field {

	field := &graphql.Field{
		Type: graphql.NewList(&salesLineType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			salesOrder, ok := p.Source.(map[string]interface{})

			if ok == true {
				p.Args["key"] = "Document_No"
				p.Args["value"] = salesOrder["No"]

				_salesLine.Object.Args = p.Args
				return _salesLine.Filter()
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
