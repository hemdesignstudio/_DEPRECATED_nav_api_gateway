// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

/*
Package salesinvoice implements a simple package for handling all graphql
operations related to Microsoft Navision SalesInvoice page.

Package has a type "SalesInvoice" where all the fields related to SalesInvoice are defined.
	'''
	type SalesInvoice struct {
		No                       string                `json:"No"`
		SellToCustomerNo         string                `json:"Sell_to_Customer_No"`
		SellToContactNo          string                `json:"Sell_to_Contact_No"`
		...
	}
	'''

GraphQl Object Type along with its fields, arguments and attributes are generated
from the SalesInvoice type when "CreateType" method is invoked.
*/
package salesinvoice

import (
	"github.com/graphql-go/graphql"
	"github.com/hem-nav-gateway/config"
	"github.com/hem-nav-gateway/salesline"
	"github.com/hem-nav-gateway/types"
)

// Microsoft Navision endpoint path for SalesInvoice page
var endpoint = config.SalesInvoiceEndpoint

/*
Response is utilized as Microsoft Navision returns a list of objects
when requesting SalesInvoice, It is utilized for JSON decoding

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
	Value []SalesInvoice `json:"value"`
}

type SalesInvoice struct {
	No                       string                `json:"No"`
	SellToCustomerNo         string                `json:"Sell_to_Customer_No"`
	SellToContactNo          string                `json:"Sell_to_Contact_No"`
	SellToCustomerName       string                `json:"Sell_to_Customer_Name"`
	SellToAddress            string                `json:"Sell_to_Address"`
	SellToAddress2           string                `json:"Sell_to_Address_2"`
	SellToPostCode           string                `json:"Sell_to_Post_Code"`
	SellToCity               string                `json:"Sell_to_City"`
	SellToContact            string                `json:"Sell_to_Contact"`
	PEBNoteOfGoods           string                `json:"PEB_Note_of_Goods"`
	PostingDate              string                `json:"Posting_Date"`
	DocumentDate             string                `json:"Document_Date"`
	SalespersonCode          string                `json:"Salesperson_Code"`
	ResponsibilityCenter     string                `json:"Responsibility_Center"`
	BillToContactNo          string                `json:"Bill_to_Contact_No"`
	BillToName               string                `json:"Bill_to_Name"`
	BillToAddress            string                `json:"Bill_to_Address"`
	BillToAddress2           string                `json:"Bill_to_Address_2"`
	BillToPostCode           string                `json:"Bill_to_Post_Code"`
	BillToCity               string                `json:"Bill_to_City"`
	BillToContact            string                `json:"Bill_to_Contact"`
	ShortcutDimension1Code   string                `json:"Shortcut_Dimension_1_Code"`
	ShortcutDimension2Code   string                `json:"Shortcut_Dimension_2_Code"`
	ShipToCode               string                `json:"Ship_to_Code"`
	ShipToName               string                `json:"Ship_to_Name"`
	ShipToAddress            string                `json:"Ship_to_Address"`
	ShipToAddress2           string                `json:"Ship_to_Address_2"`
	ShipToPostCode           string                `json:"Ship_to_Post_Code"`
	ShipToCity               string                `json:"Ship_to_City"`
	ShipToContact            string                `json:"Ship_to_Contact"`
	LocationCode             string                `json:"Location_Code"`
	ShipmentMethodCode       string                `json:"Shipment_Method_Code"`
	ShippingAgentCode        string                `json:"Shipping_Agent_Code"`
	PackageTrackingNo        string                `json:"Package_Tracking_No"`
	ShipmentDate             string                `json:"Shipment_Date"`
	YourReference            string                `json:"Your_Reference"`
	IncomingDocumentEntryNo  int                   `json:"Incoming_Document_Entry_No"`
	ExternalDocumentNo       string                `json:"External_Document_No"`
	CampaignNo               string                `json:"Campaign_No"`
	AssignedUserID           string                `json:"Assigned_User_ID"`
	JobQueueStatus           string                `json:"Job_Queue_Status"`
	Status                   string                `json:"Status"`
	PaymentTermsCode         string                `json:"Payment_Terms_Code"`
	DueDate                  string                `json:"Due_Date"`
	PaymentDiscountPercent   float64               `json:"Payment_Discount_Percent"`
	PmtDiscountDate          string                `json:"Pmt_Discount_Date"`
	PaymentMethodCode        string                `json:"Payment_Method_Code"`
	DirectDebitMandateID     string                `json:"Direct_Debit_Mandate_ID"`
	PricesIncludingVAT       bool                  `json:"Prices_Including_VAT"`
	VATBusPostingGroup       string                `json:"VAT_Bus_Posting_Group"`
	CurrencyCode             string                `json:"Currency_Code"`
	EU3PartyTrade            bool                  `json:"EU_3_Party_Trade"`
	TransactionType          string                `json:"Transaction_Type"`
	TransactionSpecification string                `json:"Transaction_Specification"`
	TransportMethod          string                `json:"Transport_Method"`
	ExitPoint                string                `json:"Exit_Point"`
	Area                     string                `json:"Area"`
	SalesLines               []salesline.SalesLine `json:"Sales_Lines"`
}

/*
CreateType function creates a GraphQl Object Type from the 'SalesInvoice' type.

example of GraphQl Object

	'''
	graphql.NewObject(graphql.ObjectConfig{
			Name: "SalesInvoice",
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
	return types.GenerateGraphQlType("SalesInvoice", SalesInvoice{}, extraFields())
}

/*
CreateArgs function creates a GraphQl Object Type from the 'SalesInvoice'

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
some arguments are required and hence in the SalesInvoice type,
tags can be noticed

example of required fields

	No	string `json:"No" required:"true"`

and this will be translated to

	"No":	&graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},


The returned GraphQl arguments will be used as a part of the main mutation
*/
func CreateArgs() map[string]*graphql.ArgumentConfig {
	return types.GenerateGraphQlArgs(SalesInvoice{}, nil)
}
