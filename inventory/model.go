// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

/*
Package inventory implements a simple package for handling all graphql
operations related to Microsoft Navision Inventory Code unit.

Package has a type "Inventory" where all the fields related to Inventory are defined.

	Example:
		'''
		type Response struct {
			Body struct {
				GetInventoryResult struct {
					RetValues struct {
						InventoryLine []Model `xml:"InventoryLine"`
			...
		}
		'''


GraphQl Object Type along with its fields, arguments and attributes are generated
from the Inventory type when "CreateType" method is invoked.
*/
package inventory

import (
	"github.com/graphql-go/graphql"
	"github.com/hem-nav-gateway/config"
	"github.com/hem-nav-gateway/types"
)

const (
	HTTP_METHOD = "POST"
	SOAP_ACTION = "urn:GetInventory"
)

var endpoint = config.InventoryEndpoint
var baseUrl = config.SoapBaseUrl

/*Response is utilized as Microsoft Navision returns a list of objects
when requesting Inventory List, It is utilized for XML decoding

Example response from Navision


Example:
	'''
	<Soap:Envelope xmlns:Soap="http://schemas.xmlsoap.org/soap/envelope/">
		<Soap:Body>
			<GetInventory_Result xmlns="urn:microsoft-dynamics-schemas/codeunit/GetInventory">
				<retValues>
						<InventoryLine xmlns="urn:microsoft-dynamics-nav/xmlports/x50000">
						<ItemNo>10005</ItemNo>
						<Description>Hai Chair Mosaic Charcoal</Description>
						<ReceiptDate>2019-03-08</ReceiptDate>
						<AvailableDate>25.00</AvailableDate>
					</InventoryLine>
				</retValues>
			</GetInventory_Result>
		</Soap:Body>
	</Soap:Envelope>
	'''

*/
type Response struct {
	Body struct {
		GetInventoryResult struct {
			RetValues struct {
				InventoryLine []Model `xml:"InventoryLine"`
			} `xml:"retValues"`
		} `xml:"GetInventory_Result"`
	}
}

type Model struct {
	ItemNo        string `xml:"ItemNo" json:"ItemNo"`
	Description   string `xml:"Description" json:"Description"`
	ReceiptDate   string `xml:"ReceiptDate" json:"ReceiptDate"`
	AvailableDate string `xml:"AvailableDate" json:"AvailableDate"`
}

/*
CreateType function creates a GraphQl Object Type from the 'Inventory' type.

Example of GraphQl Object

	Example:
		'''
		graphql.NewObject(graphql.ObjectConfig{
				Name: "Inventory",
				Fields: graphql.Fields{
					"ItemNo":		&graphql.Field{Type: graphql.String},
					"Description":		&graphql.Field{Type: graphql.String},
					"ReceiptDate":		&graphql.Field{Type: graphql.String},
					...
				},
			})
		'''

GraphQl Object is a map[string]*graphql.Field

The returned GraphQl Object Type will be used as a part of the main query
*/
func createType() *graphql.Object {
	return types.GenerateGraphQlType("Inventory", Model{}, nil)
}

/*
CreateArgs function creates a GraphQl Object Type from the 'Inventory'

Example of GraphQl Argument Object

	Example:
		'''
		map[string]*graphql.ArgumentConfig{
			"ItemNo":		&graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"Description":		&graphql.ArgumentConfig{Type: graphql.String},
			"ReceiptDate":		&graphql.ArgumentConfig{Type: graphql.String},
			...
		}
		'''

	Hint:
		Arguments are used to create or update entities,
		some arguments are required and hence in the Inventory type,
		tags can be noticed

Example of required fields

	No	string `json:"No" required:"true"`

and this will be translated to

	"No":	&graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},


The returned GraphQl arguments will be used as a part of the main mutation
*/
func createArgs() map[string]*graphql.ArgumentConfig {
	return types.GenerateGraphQlArgs(Model{}, nil)
}
