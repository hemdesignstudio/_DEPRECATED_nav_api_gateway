package item

import "github.com/graphql-go/graphql"

func CreateType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "ItemCard",
		Fields: graphql.Fields{
			"No":                            &graphql.Field{Type: graphql.String},
			"Description":                   &graphql.Field{Type: graphql.String},
			"Base_Unit_of_Measure":          &graphql.Field{Type: graphql.String},
			"Assembly_BOM":                  &graphql.Field{Type: graphql.Boolean},
			"Item_Category_Code":            &graphql.Field{Type: graphql.String},
			"Product_Group_Code":            &graphql.Field{Type: graphql.String},
			"Inventory":                     &graphql.Field{Type: graphql.Int},
			"Qty_on_Purch_Order":            &graphql.Field{Type: graphql.Int},
			"Qty_on_Sales_Order":            &graphql.Field{Type: graphql.Int},
			"Last_Date_Modified":            &graphql.Field{Type: graphql.String},
			"Freight_Type":                  &graphql.Field{Type: graphql.String},
			"Assembly_Policy":               &graphql.Field{Type: graphql.String},
			"Country_Region_of_Origin_Code": &graphql.Field{Type: graphql.String},
			"Net_Weight":                    &graphql.Field{Type: graphql.Float},
			"Gross_Weight":                  &graphql.Field{Type: graphql.Float},
			"Unit_Volume":                   &graphql.Field{Type: graphql.Float},
			"Length":                        &graphql.Field{Type: graphql.Float},
			"Width":                         &graphql.Field{Type: graphql.Float},
			"Height":                        &graphql.Field{Type: graphql.Float},
			"Designer":                      &graphql.Field{Type: graphql.String},
			"Web_Status":                    &graphql.Field{Type: graphql.String},
		},
	})
}
