package customer

import "github.com/graphql-go/graphql"

func CreateType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "CustomerCard",
		Fields: graphql.Fields{
			"No":                      &graphql.Field{Type: graphql.String},
			"Name":                    &graphql.Field{Type: graphql.String},
			"Address":                 &graphql.Field{Type: graphql.String},
			"Address_2":               &graphql.Field{Type: graphql.String},
			"Post_Code":               &graphql.Field{Type: graphql.String},
			"City":                    &graphql.Field{Type: graphql.String},
			"Country_Region_Code":     &graphql.Field{Type: graphql.String},
			"Phone_No":                &graphql.Field{Type: graphql.String},
			"Contact":                 &graphql.Field{Type: graphql.String},
			"VAT_Registration_No":     &graphql.Field{Type: graphql.String},
			"Customer_Posting_Group":  &graphql.Field{Type: graphql.String},
			"Gen_Bus_Posting_Group":   &graphql.Field{Type: graphql.String},
			"VAT_Bus_Posting_Group":   &graphql.Field{Type: graphql.String},
			"Global_Dimension_2_Code": &graphql.Field{Type: graphql.String},
			"Customer_Price_Group":    &graphql.Field{Type: graphql.String},
			"Customer_Disc_Group":     &graphql.Field{Type: graphql.String},
			"Payment_Terms_Code":      &graphql.Field{Type: graphql.String},
			"Currency_Code":           &graphql.Field{Type: graphql.String},
			"Language_Code":           &graphql.Field{Type: graphql.String},
			"Web_E_Mail":              &graphql.Field{Type: graphql.String},
			"Web_Customer":            &graphql.Field{Type: graphql.Boolean},
		},
	})
}
