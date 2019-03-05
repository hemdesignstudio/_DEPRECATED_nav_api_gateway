// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

/*
Package postship implements a simple package for handling all graphql
operations related to Microsoft Navision PostedSalesShipment page.

Package has a type "PostShip" where all the fields related to WebSalesShipment page are defined.

		type PostShip struct {
			No                       string `json:"No" required:"true"`
			SellToCustomerNo         string `json:"Sell_to_Customer_No"`
			SellToContactNo          string `json:"Sell_to_Contact_No"`
			...
		}



GraphQl Object Type along with its fields, arguments and attributes are generated
from the PostShip type when "CreateType" method is invoked.
*/
package postship

import (
	"github.com/graphql-go/graphql"
	"github.com/hem-nav-gateway/config"
	"github.com/hem-nav-gateway/types"
)

//Microsoft Navision endpoint path for WebSalesShipment page
var endpoint = config.PostShipEndpoint

/*
Response is utilized as Microsoft Navision returns a list of objects
when requesting WebSalesShipment, It is utilized for JSON decoding

example response from Navision

		{
			"value": [
				{
					"No": "102001",
					"Sell_to_Customer_No": "405124",
					"Sell_to_Contact_No": "",
					"Sell_to_Customer_Name": "Goodform",
					...
				},
				{
					"No": "102001",
					"Sell_to_Customer_No": "405124",
					"Sell_to_Contact_No": "",
					"Sell_to_Customer_Name": "Goodform",
					...
				},
				{
				...

				},
		}

*/
type Response struct {
	Value []PostShip `json:"value"`
}

type PostShip struct {
	No                       string `json:"No" required:"true"`
	SellToCustomerNo         string `json:"Sell_to_Customer_No"`
	SellToContactNo          string `json:"Sell_to_Contact_No"`
	SellToCustomerName       string `json:"Sell_to_Customer_Name"`
	SellToAddress            string `json:"Sell_to_Address"`
	SellToAddress2           string `json:"Sell_to_Address_2"`
	SellToPostCode           string `json:"Sell_to_Post_Code"`
	SellToCity               string `json:"Sell_to_City"`
	SellToContact            string `json:"Sell_to_Contact"`
	NoteOfGoods              string `json:"PEB_Note_of_Goods"`
	NoPrinted                int    `json:"No_Printed"`
	PostingDate              string `json:"Posting_Date"`
	DocumentDate             string `json:"Document_Date"`
	RequestedDeliveryDate    string `json:"Requested_Delivery_Date"`
	PromisedDeliveryDate     string `json:"Promised_Delivery_Date"`
	QuoteNo                  string `json:"Quote_No"`
	OrderNo                  string `json:"Order_No"`
	ExternalDocumentNo       string `json:"External_Document_No"`
	SalespersonCode          string `json:"Salesperson_Code"`
	ResponsibilityCenter     string `json:"Responsibility_Center"`
	BillToCustomerNo         string `json:"Bill_to_Customer_No"`
	BilltoContactNo          string `json:"Bill_to_Contact_No"`
	BillToName               string `json:"Bill_to_Name"`
	BillToAddress            string `json:"Bill_to_Address"`
	BillToAddress2           string `json:"Bill_to_Address_2"`
	BillToPostCode           string `json:"Bill_to_Post_Code"`
	BillToCity               string `json:"Bill_to_City"`
	BillToContact            string `json:"Bill_to_Contact"`
	ShortcutDimension1Code   string `json:"Shortcut_Dimension_1_Code"`
	ShortcutDimension2Code   string `json:"Shortcut_Dimension_2_Code"`
	ShipToCode               string `json:"Ship_to_Code"`
	ShipToName               string `json:"Ship_to_Name"`
	ShipToAddress            string `json:"Ship_to_Address"`
	ShipToAddress2           string `json:"Ship_to_Address_2"`
	ShipToPostCode           string `json:"Ship_to_Post_Code"`
	ShipToCity               string `json:"Ship_to_City"`
	ShipToCountryRegionCode  string `json:"Ship_to_Country_Region_Code"`
	ShipToContact            string `json:"Ship_to_Contact"`
	LocationCode             string `json:"Location_Code"`
	OutboundWhseHandlingTime string `json:"Outbound_Whse_Handling_Time"`
	ShippingTime             string `json:"Shipping_Time"`
	ShipmentMethodCode       string `json:"Shipment_Method_Code"`
	ShippingAgentCode        string `json:"Shipping_Agent_Code"`
	ShippingAgentServiceCode string `json:"Shipping_Agent_Service_Code"`
	PackageTrackingNo        string `json:"Package_Tracking_No"`
	ShipmentDate             string `json:"Shipment_Date"`
}

/*
CreateType function creates a GraphQl Object Type from the 'PostShip' type.

example of GraphQl Object

		graphql.NewObject(graphql.ObjectConfig{
				Name: "PostShip",
				Fields: graphql.Fields{
					"No":					&graphql.Field{Type: graphql.String},
					"Sell_to_Customer_No":	&graphql.Field{Type: graphql.String},
					"Sell_to_Contact_No":	&graphql.Field{Type: graphql.String},
					...
				},
			})

GraphQl Object is a map[string]*graphql.Field

The returned GraphQl Object Type will be used as a part of the main query
*/
func CreateType() *graphql.Object {
	return types.GenerateGraphQlType("PostShip", PostShip{}, nil)
}
