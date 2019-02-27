package salesinvoice

import (
	"github.com/graphql-go/graphql"
	"github.com/nav-api-gateway/config"
	"github.com/nav-api-gateway/salesline"
	"github.com/nav-api-gateway/types"
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

func CreateType() *graphql.Object {
	return types.GenerateGraphQlType("SalesInvoice", SalesInvoice{}, extraFields())
}

func CreateArgs() map[string]*graphql.ArgumentConfig {
	return types.GenerateGraphQlArgs(SalesInvoice{}, nil)
}
