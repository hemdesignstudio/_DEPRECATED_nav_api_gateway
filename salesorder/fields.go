package salesorder

import (
	"github.com/graphql-go/graphql"
	"github.com/nav-api-gateway/config"
	"github.com/nav-api-gateway/salesline"
	"log"
)

var types = map[string]*graphql.Object{
	"salesLine": salesline.CreateType("Order_SalesLines"),
}

func getSalesLinesFields() *graphql.Field {
	field := &graphql.Field{
		Type: graphql.NewList(types["salesLine"]),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			log.Printf("fetching sales lines of Company: %s which are related to Sales orders", config.CompanyName)
			salesOrder, _ := p.Source.(SalesOrder)
			p.Args["key"] = "Document_No"
			p.Args["value"] = salesOrder.No
			return salesline.Filter(p.Args)
		},
	}
	return field
}

func CreateType() *graphql.Object {
	fields := graphql.Fields{
		"No":                              &graphql.Field{Type: graphql.String},
		"Sell_to_Customer_No":             &graphql.Field{Type: graphql.String},
		"Sell_to_Customer_Name":           &graphql.Field{Type: graphql.String},
		"Sell_to_Address":                 &graphql.Field{Type: graphql.String},
		"Sell_to_Post_Code":               &graphql.Field{Type: graphql.String},
		"Sell_to_City":                    &graphql.Field{Type: graphql.String},
		"Sell_to_Contact":                 &graphql.Field{Type: graphql.String},
		"No_of_Archived_Versions":         &graphql.Field{Type: graphql.Int},
		"Customer_No_Web":                 &graphql.Field{Type: graphql.String},
		"Posting_Date":                    &graphql.Field{Type: graphql.String},
		"Order_Date":                      &graphql.Field{Type: graphql.String},
		"Document_Date":                   &graphql.Field{Type: graphql.String},
		"Requested_Delivery_Date":         &graphql.Field{Type: graphql.String},
		"Promised_Delivery_Date":          &graphql.Field{Type: graphql.String},
		"External_Document_No":            &graphql.Field{Type: graphql.String},
		"Salesperson_Code":                &graphql.Field{Type: graphql.String},
		"Assigned_User_ID":                &graphql.Field{Type: graphql.String},
		"Job_Queue_Status":                &graphql.Field{Type: graphql.String},
		"Status":                          &graphql.Field{Type: graphql.String},
		"Whs_Shipment_Lines_Exists":       &graphql.Field{Type: graphql.Boolean},
		"Bill_to_Customer_No":             &graphql.Field{Type: graphql.String},
		"Bill_to_Name":                    &graphql.Field{Type: graphql.String},
		"Bill_to_Address":                 &graphql.Field{Type: graphql.String},
		"Bill_to_Post_Code":               &graphql.Field{Type: graphql.String},
		"Bill_to_City":                    &graphql.Field{Type: graphql.String},
		"Shortcut_Dimension_2_Code":       &graphql.Field{Type: graphql.String},
		"Due_Date":                        &graphql.Field{Type: graphql.String},
		"Payment_Discount_Percent":        &graphql.Field{Type: graphql.Float},
		"Pmt_Discount_Date":               &graphql.Field{Type: graphql.String},
		"Payment_Method_Code":             &graphql.Field{Type: graphql.String},
		"Prices_Including_VAT":            &graphql.Field{Type: graphql.Boolean},
		"VAT_Bus_Posting_Group":           &graphql.Field{Type: graphql.String},
		"Ship_to_Name":                    &graphql.Field{Type: graphql.String},
		"Ship_to_Address":                 &graphql.Field{Type: graphql.String},
		"Ship_to_Address_2":               &graphql.Field{Type: graphql.String},
		"Ship_to_Post_Code":               &graphql.Field{Type: graphql.String},
		"Ship_to_City":                    &graphql.Field{Type: graphql.String},
		"Ship_to_Country_Region_Code":     &graphql.Field{Type: graphql.String},
		"Location_Code":                   &graphql.Field{Type: graphql.String},
		"Late_Order_Shipping":             &graphql.Field{Type: graphql.Boolean},
		"Shipment_Date":                   &graphql.Field{Type: graphql.String},
		"Shipping_Advice":                 &graphql.Field{Type: graphql.String},
		"Currency_Code":                   &graphql.Field{Type: graphql.String},
		"EU_3_Party_Trade":                &graphql.Field{Type: graphql.Boolean},
		"Prepayment_Percent":              &graphql.Field{Type: graphql.Float},
		"Compress_Prepayment":             &graphql.Field{Type: graphql.Boolean},
		"Prepayment_Due_Date":             &graphql.Field{Type: graphql.String},
		"Prepmt_Payment_Discount_Percent": &graphql.Field{Type: graphql.Float},
		"Prepmt_Pmt_Discount_Date":        &graphql.Field{Type: graphql.String},
		"Sales_Lines":                     getSalesLinesFields(),
	}
	return graphql.NewObject(graphql.ObjectConfig{
		Name:   "SalesOrder",
		Fields: fields,
	})
}
