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

var itemCardArgs = map[string]*graphql.ArgumentConfig{
	"No":                            {Type: graphql.NewNonNull(graphql.String)},
	"Description":                   {Type: graphql.String},
	"Base_Unit_of_Measure":          {Type: graphql.String},
	"Assembly_BOM":                  {Type: graphql.Boolean},
	"Item_Category_Code":            {Type: graphql.String},
	"Product_Group_Code":            {Type: graphql.String},
	"Inventory":                     {Type: graphql.Int},
	"Qty_on_Purch_Order":            {Type: graphql.Int},
	"Qty_on_Sales_Order":            {Type: graphql.Int},
	"Last_Date_Modified":            {Type: graphql.String},
	"Freight_Type":                  {Type: graphql.String},
	"Assembly_Policy":               {Type: graphql.String},
	"Country_Region_of_Origin_Code": {Type: graphql.String},
	"Net_Weight":                    {Type: graphql.Float},
	"Gross_Weight":                  {Type: graphql.Float},
	"Unit_Volume":                   {Type: graphql.Float},
	"Length":                        {Type: graphql.Float},
	"Width":                         {Type: graphql.Float},
	"Height":                        {Type: graphql.Float},
	"Designer":                      {Type: graphql.String},
	"Web_Status":                    {Type: graphql.String},
}
