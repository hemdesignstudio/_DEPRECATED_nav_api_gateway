package salesinvoice

import (
	"github.com/graphql-go/graphql"
	"github.com/nav-api-gateway/config"
	"github.com/nav-api-gateway/salesline"
	"log"
)

var types = map[string]*graphql.Object{
	"salesLine": salesline.CreateType("Invoice_SalesLines"),
}

func getSalesLinesFields() *graphql.Field {
	field := &graphql.Field{
		Type: graphql.NewList(types["salesLine"]),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			log.Printf("fetching sales lines of company: %s which are related to Sales Invoices", config.CompanyName)
			salesInvoice, _ := p.Source.(SalesInvoice)
			p.Args["key"] = "Document_No"
			p.Args["value"] = salesInvoice.No
			return salesline.Filter(p.Args)
		},
	}
	return field
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
		"PEB_Note_of_Goods":          &graphql.Field{Type: graphql.String},
		"Posting_Date":               &graphql.Field{Type: graphql.String},
		"Document_Date":              &graphql.Field{Type: graphql.String},
		"Salesperson_Code":           &graphql.Field{Type: graphql.String},
		"Responsibility_Center":      &graphql.Field{Type: graphql.String},
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
