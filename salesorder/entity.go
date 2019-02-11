package salesorder

import (
	"github.com/graphql-go/graphql"
	"github.com/nav-api-gateway/config"
	"github.com/nav-api-gateway/salesline"
	"github.com/nav-api-gateway/types"
)

var endpoint = config.SalesOrderEndpoint
var companyName = config.CompanyName

type Response struct {
	Value []SalesOrder `json:"value"`
}

type SalesOrder struct {
	No                           string                `json:"No"`
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
	BillToCustomerNo             string                `json:"Bill_to_Customer_No"`
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

func CreateType() *graphql.Object {
	return types.GenerateGraphQlType("SalesOrder", SalesOrder{}, extraFields())
}
