package company

import "github.com/graphql-go/graphql"

var filterArgs = map[string]*graphql.ArgumentConfig{
	"key":   {Type: graphql.String},
	"value": {Type: graphql.String},
}

var customerCardArgs = map[string]*graphql.ArgumentConfig{
	"No":                      {Type: graphql.NewNonNull(graphql.String)},
	"Name":                    {Type: graphql.String},
	"Address":                 {Type: graphql.String},
	"Address_2":               {Type: graphql.String},
	"Post_Code":               {Type: graphql.String},
	"City":                    {Type: graphql.String},
	"Country_Region_Code":     {Type: graphql.String},
	"Phone_No":                {Type: graphql.String},
	"Contact":                 {Type: graphql.String},
	"VAT_Registration_No":     {Type: graphql.String},
	"Customer_Posting_Group":  {Type: graphql.String},
	"Gen_Bus_Posting_Group":   {Type: graphql.String},
	"VAT_Bus_Posting_Group":   {Type: graphql.String},
	"Global_Dimension_2_Code": {Type: graphql.String},
	"Customer_Price_Group":    {Type: graphql.String},
	"Customer_Disc_Group":     {Type: graphql.String},
	"Payment_Terms_Code":      {Type: graphql.String},
	"Currency_Code":           {Type: graphql.String},
	"Language_Code":           {Type: graphql.String},
	"Web_E_Mail":              {Type: graphql.String},
	"Web_Customer":            {Type: graphql.Boolean},
}
