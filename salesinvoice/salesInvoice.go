package salesinvoice

import (
	"encoding/json"
	"github.com/graphql-go/graphql"
	"github.com/nav-api-gateway/config"
	errhandler "github.com/nav-api-gateway/error"
	"github.com/nav-api-gateway/request"
	"github.com/nav-api-gateway/salesline"
)

var endpoint = config.SalesInvoiceEndpoint
var companyName = config.CompanyName

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
	NoteOfGoods              string                `json:"Note_of_Goods"`
	PostingDate              string                `json:"Posting_Date"`
	DocumentDate             string                `json:"Document_Date"`
	SalespersonCode          string                `json:"Salesperson_Code"`
	ResponsibilityCenter     string                `json:"Responsibility_Center"`
	BillToCustomerNo         string                `json:"Bill_to_Customer_No"`
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
	CreditCardNo             string                `json:"Credit_Card_No"`
	GetCreditCardNumber      string                `json:"Get_Credit_Card_Number"`
	CurrencyCode             string                `json:"Currency_Code"`
	EU3PartyTrade            bool                  `json:"EU_3_Party_Trade"`
	TransactionType          string                `json:"Transaction_Type"`
	TransactionSpecification string                `json:"Transaction_Specification"`
	TransportMethod          string                `json:"Transport_Method"`
	ExitPoint                string                `json:"Exit_Point"`
	Area                     string                `json:"Area"`
	SalesLines               []salesline.SalesLine `json:"Sales_Lines"`
}

func CreateType() *graphql.Object {
	fields := graphql.Fields{
		"No":                         &graphql.Field{Type: graphql.String},
		"Sell_to_Customer_No":        &graphql.Field{Type: graphql.String},
		"Sell_to_Contact_No":         &graphql.Field{Type: graphql.String},
		"Sell_to_Customer_Name":      &graphql.Field{Type: graphql.String},
		"Sell_to_Address":            &graphql.Field{Type: graphql.String},
		"Sell_to_Address_2":          &graphql.Field{Type: graphql.String},
		"Sell_to_Post_Code":          &graphql.Field{Type: graphql.String},
		"Sell_to_City":               &graphql.Field{Type: graphql.String},
		"Sell_to_Contact":            &graphql.Field{Type: graphql.String},
		"Note_of_Goods":              &graphql.Field{Type: graphql.String},
		"Posting_Date":               &graphql.Field{Type: graphql.String},
		"Document_Date":              &graphql.Field{Type: graphql.String},
		"Salesperson_Code":           &graphql.Field{Type: graphql.String},
		"Responsibility_Center":      &graphql.Field{Type: graphql.String},
		"Bill_to_Customer_No":        &graphql.Field{Type: graphql.String},
		"Bill_to_Contact_No":         &graphql.Field{Type: graphql.String},
		"Bill_to_Name":               &graphql.Field{Type: graphql.String},
		"Bill_to_Address":            &graphql.Field{Type: graphql.String},
		"Bill_to_Address_2":          &graphql.Field{Type: graphql.String},
		"Bill_to_Post_Code":          &graphql.Field{Type: graphql.String},
		"Bill_to_City":               &graphql.Field{Type: graphql.String},
		"Bill_to_Contact":            &graphql.Field{Type: graphql.String},
		"Shortcut_Dimension_1_Code":  &graphql.Field{Type: graphql.String},
		"Shortcut_Dimension_2_Code":  &graphql.Field{Type: graphql.String},
		"Ship_to_Code":               &graphql.Field{Type: graphql.String},
		"Ship_to_Name":               &graphql.Field{Type: graphql.String},
		"Ship_to_Address":            &graphql.Field{Type: graphql.String},
		"Ship_to_Address_2":          &graphql.Field{Type: graphql.String},
		"Ship_to_Post_Code":          &graphql.Field{Type: graphql.String},
		"Ship_to_City":               &graphql.Field{Type: graphql.String},
		"Ship_to_Contact":            &graphql.Field{Type: graphql.String},
		"Location_Code":              &graphql.Field{Type: graphql.String},
		"Shipment_Method_Code":       &graphql.Field{Type: graphql.String},
		"Shipping_Agent_Code":        &graphql.Field{Type: graphql.String},
		"Package_Tracking_No":        &graphql.Field{Type: graphql.String},
		"Shipment_Date":              &graphql.Field{Type: graphql.String},
		"Your_Reference":             &graphql.Field{Type: graphql.String},
		"Incoming_Document_Entry_No": &graphql.Field{Type: graphql.Int},
		"External_Document_No":       &graphql.Field{Type: graphql.String},
		"Campaign_No":                &graphql.Field{Type: graphql.String},
		"Assigned_User_ID":           &graphql.Field{Type: graphql.String},
		"Job_Queue_Status":           &graphql.Field{Type: graphql.String},
		"Status":                     &graphql.Field{Type: graphql.String},
		"Payment_Terms_Code":         &graphql.Field{Type: graphql.String},
		"Due_Date":                   &graphql.Field{Type: graphql.String},
		"Payment_Discount_Percent":   &graphql.Field{Type: graphql.Float},
		"Pmt_Discount_Date":          &graphql.Field{Type: graphql.String},
		"Payment_Method_Code":        &graphql.Field{Type: graphql.String},
		"Direct_Debit_Mandate_ID":    &graphql.Field{Type: graphql.String},
		"Prices_Including_VAT":       &graphql.Field{Type: graphql.Boolean},
		"VAT_Bus_Posting_Group":      &graphql.Field{Type: graphql.String},
		"Credit_Card_No":             &graphql.Field{Type: graphql.String},
		"Get_Credit_Card_Number":     &graphql.Field{Type: graphql.String},
		"Currency_Code":              &graphql.Field{Type: graphql.String},
		"EU_3_Party_Trade":           &graphql.Field{Type: graphql.Boolean},
		"Transaction_Type":           &graphql.Field{Type: graphql.String},
		"Transaction_Specification":  &graphql.Field{Type: graphql.String},
		"Transport_Method":           &graphql.Field{Type: graphql.String},
		"Exit_Point":                 &graphql.Field{Type: graphql.String},
		"Area":                       &graphql.Field{Type: graphql.String},
		"Sales_Lines":                getSalesLinesFields(),
	}

	return graphql.NewObject(graphql.ObjectConfig{
		Name:   "SalesInvoice",
		Fields: fields,
	})
}

func GetAll() ([]SalesInvoice, error) {
	resultByte := request.GetAll(config.CompanyName, endpoint)
	res := Response{}
	err := json.Unmarshal(resultByte, &res)
	if err != nil {
		return nil, errhandler.CouldNotUnmarshalData()
	}
	return res.Value, nil
}

func Filter(args map[string]interface{}) ([]SalesInvoice, error) {
	resByte, resError := request.Filter(config.CompanyName, endpoint, args)
	if resError != nil {
		return nil, resError
	}
	res := Response{}
	err := json.Unmarshal(resByte, &res)
	if err != nil {
		return nil, errhandler.CouldNotUnmarshalData()
	}
	return res.Value, nil
}

func Create(args map[string]interface{}) (SalesInvoice, error) {
	body, _ := json.Marshal(args)
	resByte := request.Create(config.CompanyName, endpoint, body)
	res := SalesInvoice{}
	err := json.Unmarshal(resByte, &res)
	if err != nil {
		return res, errhandler.CouldNotUnmarshalData()
	}
	return res, nil
}

func Update(args map[string]interface{}) (SalesInvoice, error) {
	no := args["No"].(string)
	body, _ := json.Marshal(args)
	resByte := request.Update(companyName, endpoint, no, body)
	res := SalesInvoice{}
	err := json.Unmarshal(resByte, &res)
	if err != nil {
		return res, errhandler.CouldNotUnmarshalData()
	}
	return res, nil
}
