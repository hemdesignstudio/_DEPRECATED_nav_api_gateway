package postship

import (
	"github.com/nav-api-gateway/config"
)

var endpoint = config.PostShipEndpoint
var companyName = config.CompanyName

type Response struct {
	Value []PostShip `json:"value"`
}

type PostShip struct {
	No                       string `json:"No"`
	SellToCustomerNo         string `json:"Sell_to_Customer_No"`
	SellToContactNo          string `json:"Sell_to_Contact_No"`
	SellToCustomerName       string `json:"Sell_to_Customer_Name"`
	SellToAddress            string `json:"Sell_to_Address"`
	SellToAddress2           string `json:"Sell_to_Address2"`
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
	BillToAddress2           string `json:"Bill_to_Address2"`
	BillToPostCode           string `json:"Bill_to_Post_Code"`
	BillToCity               string `json:"Bill_to_City"`
	BillToContact            string `json:"Bill_to_Contact"`
	ShortcutDimension1Code   string `json:"Shortcut_Dimension_1_Code"`
	ShortcutDimension2Code   string `json:"Shortcut_Dimension_2_Code"`
	ShipToCode               string `json:"Ship_to_Code"`
	ShipToName               string `json:"Ship_to_Name"`
	ShipToAddress            string `json:"Ship_to_Address"`
	ShipToAddress2           string `json:"Ship_to_Address2"`
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
