package postship

import "github.com/graphql-go/graphql"

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
