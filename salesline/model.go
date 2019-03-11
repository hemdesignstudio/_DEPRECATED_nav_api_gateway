// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

/*
Package salesline implements a simple package for handling all graphql
operations related to Microsoft Navision SalesLine page.

Package has a type "SalesLine" where all the fields related to SalesLine are defined.

	Example
		'''
		type SalesLine struct {
			No                  string  `json:"No"`
			DocumentNo          string  `json:"Document_No" required:"true"`
			DocumentType        string  `json:"Document_Type" required:"true"`
			LineNo              int     `json:"Line_No" required:"true"
			...
		}
		'''

GraphQl Object Type along with its fields, arguments and attributes are generated
from the SalesLine type when "CreateType" method is invoked.
*/
package salesline

import (
	"github.com/graphql-go/graphql"
	"github.com/hem-nav-gateway/config"
	"github.com/hem-nav-gateway/types"
)

// endpoint refers to Microsoft Navision endpoint path for SalesLine page

var endpoint = config.SalesLineEndpoint

/*
Response is utilized as Microsoft Navision returns a list of objects
when requesting SalesLine, It is utilized for JSON decoding

Example response from Navision

	Example
		'''
		{
			"value": [
				{
					"No": "1234"
					"Document_No": "1001",
					"Document_Type": "Order",
					...
				},
				{
					"No": "2345"
					"Document_No": "1002",
					"Document_Type": "Order",
					...
				},
				{
				...

				},
		}
		'''
*/
type Response struct {
	Value []SalesLine `json:"value"`
}

type SalesLine struct {
	No                  string  `json:"No"`
	DocumentNo          string  `json:"Document_No" required:"true"`
	DocumentType        string  `json:"Document_Type" required:"true"`
	LineNo              int     `json:"Line_No" required:"true"`
	Type                string  `json:"Type"`
	Description         string  `json:"Description"`
	Reserve             string  `json:"Reserve"`
	Quantity            int     `json:"Quantity"`
	ReservedQuantity    int     `json:"Reserved_Quantity"`
	UnitOfMeasureCode   string  `json:"Unit_of_Measure_Code"`
	UnitPrice           float64 `json:"Unit_Price"`
	LineAmount          float64 `json:"Line_Amount"`
	LineDiscountPercent float64 `json:"Line_Discount_Percent"`
	LineDiscountAmount  float64 `json:"Line_Discount_Amount"`
	PrepaymentPercent   float64 `json:"Prepayment_Percent"`
	PrepmtLineAmount    float64 `json:"Prepmt_Line_Amount"`
	QtyToShip           int     `json:"Qty_to_Ship"`
	QuantityShipped     int     `json:"Quantity_Shipped"`
	QtyToInvoice        int     `json:"Qty_to_Invoice"`
	QuantityInvoiced    int     `json:"Quantity_Invoiced"`
}

/*
CreateType function creates a GraphQl Object Type from the 'SalesLine' type.

Example of GraphQl Object

	Example:
		'''
		graphql.NewObject(graphql.ObjectConfig{
				Name: "SalesLine",
				Fields: graphql.Fields{
					"No":				&graphql.Field{Type: graphql.String},
					"Document_No":			&graphql.Field{Type: graphql.String},
					"Document_Type":		&graphql.Field{Type: graphql.String},
					...
				},
			})
		'''`

GraphQl Object is a map[string]*graphql.Field

The returned GraphQl Object Type will be used as a part of the main query
*/
func CreateType(name string) *graphql.Object {
	return types.GenerateGraphQlType(name, SalesLine{}, nil)
}

/*
CreateArgs function creates a GraphQl Object Type from the 'SalesLine'

Example of GraphQl Argument Object

	Example:
		'''
		map[string]*graphql.ArgumentConfig{
			"No":				&graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"Document_No":			&graphql.ArgumentConfig{Type: graphql.String},
			"Document_Type":		&graphql.ArgumentConfig{Type: graphql.String},
			...
		}
		'''

Hint:
	Arguments are used to create or update entities,
	some arguments are required and hence in the SalesLine type,
	tags can be noticed.

Example of required fields

	No	string `json:"No" required:"true"`

and this will be translated to

	"No":	&graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},


The returned GraphQl arguments will be used as a part of the main mutation
*/
func CreateArgs() map[string]*graphql.ArgumentConfig {
	return types.GenerateGraphQlArgs(SalesLine{}, nil)
}
