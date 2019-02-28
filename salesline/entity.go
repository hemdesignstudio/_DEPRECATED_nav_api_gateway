package salesline

import (
	"github.com/graphql-go/graphql"
	"github.com/hem-nav-gateway/config"
	"github.com/hem-nav-gateway/types"
)

var endpoint = config.SalesLineEndpoint
var companyName = config.CompanyName

type Response struct {
	Value []SalesLine `json:"value"`
}

type SalesLine struct {
	No                  string  `json:"No"`
	DocumentNo          string  `json:"Document_No"`
	DocumentType        string  `json:"Document_Type"`
	LineNo              int     `json:"Line_No"`
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

func CreateType(name string) *graphql.Object {
	return types.GenerateGraphQlType(name, SalesLine{}, nil)
}

func CreateArgs() map[string]*graphql.ArgumentConfig {
	return types.GenerateGraphQlArgs(SalesLine{}, nil)
}
