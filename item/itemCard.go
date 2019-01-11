package item

import (
	"encoding/json"
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/nav-api-gateway/config"
	"github.com/nav-api-gateway/request"
)

type Response struct {
	Value []ItemCard `json:"value"`
}

type ItemCard struct {
	No                        string  `json:"No"`
	Description               string  `json:"Description"`
	BaseUnitOfMeasure         string  `json:"Base_Unit_of_Measure"`
	AssemblyBOM               bool    `json:"Assembly_BOM"`
	ItemCategoryCode          string  `json:"Item_Category_Code"`
	ProductGroupCode          string  `json:"Product_Group_Code"`
	Inventory                 int     `json:"Inventory"`
	QtyOnPurchOrder           int     `json:"Qty_on_Purch_Order"`
	QtyOnSalesOrder           int     `json:"Qty_on_Sales_Order"`
	LastDateModified          string  `json:"Last_Date_Modified"`
	FreightType               string  `json:"Freight_Type"`
	AssemblyPolicy            string  `json:"Assembly_Policy"`
	CountryRegionOfOriginCode string  `json:"Country_Region_of_Origin_Code"`
	NetWeight                 float32 `json:"Net_Weight"`
	GrossWeight               float32 `json:"Gross_Weight"`
	UnitVolume                float32 `json:"Unit_Volume"`
	Length                    float32 `json:"Length"`
	Width                     float32 `json:"Width"`
	Height                    float32 `json:"Height"`
	Designer                  string  `json:"Designer"`
	WebStatus                 string  `json:"Web_Status"`
}

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

func GetAll() ([]ItemCard, error) {
	resByte := request.GetAll(config.CompanyName, config.ItemCardEndpoint)
	res := Response{}
	err := json.Unmarshal(resByte, &res)
	if err != nil {
		return nil, errors.New("could not unmarshal data")
	}
	return res.Value, nil
}

func Filter(args map[string]interface{}) ([]ItemCard, error) {
	resByte := request.Filter(config.CompanyName, config.ItemCardEndpoint, args)
	res := Response{}
	err := json.Unmarshal(resByte, &res)
	if err != nil {
		return nil, errors.New("could not unmarshal data")
	}
	return res.Value, nil
}
