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
var _salesLine = salesline.Request{}
var salesLineType = *_salesLine.NewType("Invoice_SalesLines")

// getSalesLinesFields creates a Salesline field objects
// which will be used as a nested field for Sales Invoice
func getSalesLinesFields() *graphql.Field {

	field := &graphql.Field{
		Type: graphql.NewList(&salesLineType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			type result struct {
				data interface{}
				err  error
			}
			ch := make(chan *result, 1)
			go func() {
				defer close(ch)
				salesOrder, ok := p.Source.(map[string]interface{})

				if ok == true {
					p.Args["key"] = "Document_No"
					p.Args["value"] = salesOrder["No"]

					_salesLine.Object.Args = p.Args
					data, err := _salesLine.Filter()
					if err != nil {
						ch <- &result{data: nil, err: err}

					} else {
						ch <- &result{data: data, err: nil}
					}

				} else {
					ch <- &result{data: nil, err: nil}

				}
			}()
			return func() (interface{}, error) {
				r := <-ch
				return r.data, r.err
			}, nil
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
