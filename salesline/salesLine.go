package salesline

import (
	"encoding/json"
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/nav-api-gateway/config"
	"github.com/nav-api-gateway/request"
)

var endpoint = config.SalesLineEndpoint
var companyName = config.CompanyName

type Response struct {
	Value []SalesLine `json:"value"`
}

type SalesLine struct {
	No                  string  `json:"No"`
	DocumentNo          string  `json:"Document_No"`
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

	fields := graphql.Fields{
		"No":                    &graphql.Field{Type: graphql.String},
		"Document_No":           &graphql.Field{Type: graphql.String},
		"Type":                  &graphql.Field{Type: graphql.String},
		"Description":           &graphql.Field{Type: graphql.String},
		"Reserve":               &graphql.Field{Type: graphql.String},
		"Quantity":              &graphql.Field{Type: graphql.Int},
		"Reserved_Quantity":     &graphql.Field{Type: graphql.Int},
		"Unit_of_Measure_Code":  &graphql.Field{Type: graphql.String},
		"Unit_Price":            &graphql.Field{Type: graphql.Float},
		"Line_Amount":           &graphql.Field{Type: graphql.Float},
		"Line_Discount_Percent": &graphql.Field{Type: graphql.Float},
		"Line_Discount_Amount":  &graphql.Field{Type: graphql.Float},
		"Prepayment_Percent":    &graphql.Field{Type: graphql.Float},
		"Prepmt_Line_Amount":    &graphql.Field{Type: graphql.Float},
		"Qty_to_Ship":           &graphql.Field{Type: graphql.Int},
		"Quantity_Shipped":      &graphql.Field{Type: graphql.Int},
		"Qty_to_Invoice":        &graphql.Field{Type: graphql.Int},
		"Quantity_Invoiced":     &graphql.Field{Type: graphql.Int},
	}

	return graphql.NewObject(graphql.ObjectConfig{
		Name:   name,
		Fields: fields,
	})
}

func GetAll() ([]SalesLine, error) {
	resultByte := request.GetAll(companyName, endpoint)
	res := Response{}
	err := json.Unmarshal(resultByte, &res)
	if err != nil {
		return nil, errors.New("could not unmarshal data")
	}
	return res.Value, nil
}

func Filter(args map[string]interface{}) ([]SalesLine, error) {
	resByte := request.Filter(companyName, endpoint, args)
	res := Response{}
	err := json.Unmarshal(resByte, &res)
	if err != nil {
		return nil, errors.New("could not unmarshal data")
	}
	return res.Value, nil
}

func Create(args map[string]interface{}) (SalesLine, error) {
	body, _ := json.Marshal(args)
	resByte := request.Create(companyName, endpoint, body)
	res := SalesLine{}
	err := json.Unmarshal(resByte, &res)
	if err != nil {
		return res, errors.New("could not unmarshal data")
	}
	return res, nil
}

func Update(args map[string]interface{}) (string, error) {
	no := args["No"].(string)
	body, _ := json.Marshal(args)
	status := request.Update(companyName, endpoint, no, body)
	return status, nil
}
