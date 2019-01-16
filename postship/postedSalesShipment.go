package postship

import (
	"encoding/json"
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/nav-api-gateway/config"
	"github.com/nav-api-gateway/request"
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

func CreateType() *graphql.Object {
	fields := graphql.Fields{
		"No":                          &graphql.Field{Type: graphql.String},
		"Sell_to_Customer_No":         &graphql.Field{Type: graphql.String},
		"Sell_to_Contact_No":          &graphql.Field{Type: graphql.String},
		"Sell_to_Customer_Name":       &graphql.Field{Type: graphql.String},
		"Sell_to_Address":             &graphql.Field{Type: graphql.String},
		"Sell_to_Address2":            &graphql.Field{Type: graphql.String},
		"Sell_to_Post_Code":           &graphql.Field{Type: graphql.String},
		"Sell_to_City":                &graphql.Field{Type: graphql.String},
		"Sell_to_Contact":             &graphql.Field{Type: graphql.String},
		"PEB_Note_of_Goods":           &graphql.Field{Type: graphql.String},
		"No_Printed":                  &graphql.Field{Type: graphql.Int},
		"Posting_Date":                &graphql.Field{Type: graphql.String},
		"Document_Date":               &graphql.Field{Type: graphql.String},
		"Requested_Delivery_Date":     &graphql.Field{Type: graphql.String},
		"Promised_Delivery_Date":      &graphql.Field{Type: graphql.String},
		"Quote_No":                    &graphql.Field{Type: graphql.String},
		"Order_No":                    &graphql.Field{Type: graphql.String},
		"External_Document_No":        &graphql.Field{Type: graphql.String},
		"Salesperson_Code":            &graphql.Field{Type: graphql.String},
		"Responsibility_Center":       &graphql.Field{Type: graphql.String},
		"Bill_to_Customer_No":         &graphql.Field{Type: graphql.String},
		"Bill_to_Contact_No":          &graphql.Field{Type: graphql.String},
		"Bill_to_Name":                &graphql.Field{Type: graphql.String},
		"Bill_to_Address":             &graphql.Field{Type: graphql.String},
		"Bill_to_Address2":            &graphql.Field{Type: graphql.String},
		"Bill_to_Post_Code":           &graphql.Field{Type: graphql.String},
		"Bill_to_City":                &graphql.Field{Type: graphql.String},
		"Bill_to_Contact":             &graphql.Field{Type: graphql.String},
		"Shortcut_Dimension_1_Code":   &graphql.Field{Type: graphql.String},
		"Shortcut_Dimension_2_Code":   &graphql.Field{Type: graphql.String},
		"Ship_to_Code":                &graphql.Field{Type: graphql.String},
		"Ship_to_Name":                &graphql.Field{Type: graphql.String},
		"Ship_to_Address":             &graphql.Field{Type: graphql.String},
		"Ship_to_Address2":            &graphql.Field{Type: graphql.String},
		"Ship_to_Post_Code":           &graphql.Field{Type: graphql.String},
		"Ship_to_City":                &graphql.Field{Type: graphql.String},
		"Ship_to_Country_Region_Code": &graphql.Field{Type: graphql.String},
		"Ship_to_Contact":             &graphql.Field{Type: graphql.String},
		"Location_Code":               &graphql.Field{Type: graphql.String},
		"Outbound_Whse_Handling_Time": &graphql.Field{Type: graphql.String},
		"Shipping_Time":               &graphql.Field{Type: graphql.String},
		"Shipment_Method_Code":        &graphql.Field{Type: graphql.String},
		"Shipping_Agent_Code":         &graphql.Field{Type: graphql.String},
		"Shipping_Agent_Service_Code": &graphql.Field{Type: graphql.String},
		"Package_Tracking_No":         &graphql.Field{Type: graphql.String},
		"Shipment_Date":               &graphql.Field{Type: graphql.String},
	}
	return graphql.NewObject(graphql.ObjectConfig{
		Name:   "PostShip",
		Fields: fields,
	})
}

func GetAll() ([]PostShip, error) {
	resByte := request.GetAll(companyName, endpoint)
	res := Response{}
	err := json.Unmarshal(resByte, &res)
	if err != nil {
		return nil, errors.New("could not unmarshal data")
	}
	return res.Value, nil
}

func Filter(args map[string]interface{}) ([]PostShip, error) {
	resByte := request.Filter(companyName, endpoint, args)
	res := Response{}
	err := json.Unmarshal(resByte, &res)
	if err != nil {
		return nil, errors.New("could not unmarshal data")
	}
	return res.Value, nil
}

func Create(args map[string]interface{}) (PostShip, error) {
	body, _ := json.Marshal(args)
	resByte := request.Create(companyName, endpoint, body)
	res := PostShip{}
	err := json.Unmarshal(resByte, &res)
	if err != nil {
		return res, errors.New("could not unmarshal data")
	}
	return res, nil
}

func Update(args map[string]interface{}) (PostShip, error) {
	no := args["No"].(string)
	body, _ := json.Marshal(args)
	resByte := request.Update(companyName, endpoint, no, body)
	res := PostShip{}
	err := json.Unmarshal(resByte, &res)
	if err != nil {
		return res, errors.New("could not unmarshal data")
	}
	return res, nil
}
