// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

/*
Package SalesOrder implements a simple package for handling all graphql
operations related to Microsoft Navision SalesOrder page.

Package has a type "SalesOrder" where all the fields related to SalesOrder are defined.

	'''
	type SalesOrder struct {
		No                       string                `json:"No"`
		SellToCustomerNo         string                `json:"Sell_to_Customer_No"`
		SellToContactNo          string                `json:"Sell_to_Contact_No"`
		...
	}
	'''

GraphQl Object Type along with its fields, arguments and attributes are generated
from the SalesOrder type when "CreateType" method is invoked.
*/
package salesorder

import (
	"github.com/graphql-go/graphql"
	"github.com/hem-nav-gateway/config"
	"github.com/hem-nav-gateway/salesline"
	"github.com/hem-nav-gateway/types"
)

//endpoint refers to Microsoft Navision endpoint path for SalesOrderWs page
var endpoint = config.SalesOrderEndpoint

/*
Response is utilized as Microsoft Navision returns a list of objects
when requesting SalesOrder, It is utilized for JSON decoding

example response from Navision

	'''
	{
		"value": [
			{
				"No": "100001",
				"Sell_to_Customer_No": "206901",
				"Sell_to_Customer_Name": "Studio Feuer AB",

				...
			},
			{
				"No": "100002",
				"Sell_to_Customer_No": "WEB-10",
				"Sell_to_Customer_Name": "Web Customer Currency EUR",

				...
			},
			{
			...

			},
	}
	'''

*/
type Response struct {
	Value []SalesOrder `json:"value"`
}

type SalesOrder struct {
	No                           string                `json:"No" required:"true"`
	SellToCustomerNo             string                `json:"Sell_to_Customer_No"`
	SellToCustomerName           string                `json:"Sell_to_Customer_Name"`
	SellToAddress                string                `json:"Sell_to_Address"`
	SellToPostCode               string                `json:"Sell_to_Post_Code"`
	SellToCity                   string                `json:"Sell_to_City"`
	SellToContact                string                `json:"Sell_to_Contact"`
	NoOfArchivedVersions         int                   `json:"No_of_Archived_Versions"`
	CustomerNoWeb                string                `json:"Customer_No_Web"`
	PostingDate                  string                `json:"Posting_Date"`
	OrderDate                    string                `json:"Order_Date"`
	DocumentDate                 string                `json:"Document_Date"`
	RequestedDeliveryDate        string                `json:"Requested_Delivery_Date"`
	PromisedDeliveryDate         string                `json:"Promised_Delivery_Date"`
	ExternalDocumentNo           string                `json:"External_Document_No"`
	SalespersonCode              string                `json:"Salesperson_Code"`
	AssignedUserID               string                `json:"Assigned_User_ID"`
	JobQueueStatus               string                `json:"Job_Queue_Status"`
	Status                       string                `json:"Status"`
	WhsShipmentLinesExists       bool                  `json:"Whs_Shipment_Lines_Exists"`
	BillToName                   string                `json:"Bill_to_Name"`
	BillToAddress                string                `json:"Bill_to_Address"`
	BillToPostCode               string                `json:"Bill_to_Post_Code"`
	BillToCity                   string                `json:"Bill_to_City"`
	ShortcutDimension2Code       string                `json:"Shortcut_Dimension_2_Code"`
	DueDate                      string                `json:"Due_Date"`
	PaymentDiscountPercent       float32               `json:"Payment_Discount_Percent"`
	PmtDiscountDate              string                `json:"Pmt_Discount_Date"`
	PaymentMethodCode            string                `json:"Payment_Method_Code"`
	PricesIncludingVAT           bool                  `json:"Prices_Including_VAT"`
	VATBusPostingGroup           string                `json:"VAT_Bus_Posting_Group"`
	ShipToName                   string                `json:"Ship_to_Name"`
	ShipToAddress                string                `json:"Ship_to_Address"`
	ShipToAddress2               string                `json:"Ship_to_Address_2"`
	ShipToPostCode               string                `json:"Ship_to_Post_Code"`
	ShipToCity                   string                `json:"Ship_to_City"`
	ShipToCountryRegionCode      string                `json:"Ship_to_Country_Region_Code"`
	LocationCode                 string                `json:"Location_Code"`
	LateOrderShipping            bool                  `json:"Late_Order_Shipping"`
	ShipmentDate                 string                `json:"Shipment_Date"`
	ShippingAdvice               string                `json:"Shipping_Advice"`
	CurrencyCode                 string                `json:"Currency_Code"`
	EU3PartyTrade                bool                  `json:"EU_3_Party_Trade"`
	PrepaymentPercent            float32               `json:"Prepayment_Percent"`
	CompressPrepayment           bool                  `json:"Compress_Prepayment"`
	PrepaymentDueDate            string                `json:"Prepayment_Due_Date"`
	PrepmtPaymentDiscountPercent float32               `json:"Prepmt_Payment_Discount_Percent"`
	PrepmtPmtDiscountDate        string                `json:"Prepmt_Pmt_Discount_Date"`
	SalesLines                   []salesline.SalesLine `json:"Sales_Lines"`
}

/*
CreateType function creates a GraphQl Object Type from the 'SalesOrder' type.

example of GraphQl Object

	'''
	graphql.NewObject(graphql.ObjectConfig{
			Name: "SalesOrder",
			Fields: graphql.Fields{
				"No":				&graphql.Field{Type: graphql.String},
				"Sell_to_Customer_No":		&graphql.Field{Type: graphql.String},
				"Sell_to_Customer_Name":	&graphql.Field{Type: graphql.String},
				...
			},
		})
	'''

GraphQl Object is a map[string]*graphql.Field

The returned GraphQl Object Type will be used as a part of the main query
*/
func CreateType() *graphql.Object {
	return types.GenerateGraphQlType("SalesOrder", SalesOrder{}, extraFields())
}

/*
CreateArgs function creates a GraphQl Object Type from the 'SalesOrder'

example of GraphQl Argument Object

	'''
	map[string]*graphql.ArgumentConfig{
		"No":				&graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"Sell_to_Customer_No":		&graphql.ArgumentConfig{Type: graphql.String},
		"Sell_to_Customer_Name":	&graphql.ArgumentConfig{Type: graphql.String},
		...
	}
	'''

Hint: arguments are used to create or update entities,
some arguments are required and hence in the SalesOrder type,
tags can be noticed

example of required fields

	No	string `json:"No" required:"true"`

and this will be translated to

	"No":	&graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},


The returned GraphQl arguments will be used as a part of the main mutation
*/
func CreateArgs() map[string]*graphql.ArgumentConfig {
	return types.GenerateGraphQlArgs(SalesOrder{}, nil)
}
