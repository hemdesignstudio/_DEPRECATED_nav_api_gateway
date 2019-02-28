package item

import (
	"github.com/graphql-go/graphql"
	"github.com/nav-api-gateway/config"
	"github.com/nav-api-gateway/types"
)

var endpoint = config.ItemCardEndpoint
var companyName = config.CompanyName

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
	return types.GenerateGraphQlType("ItemCard", ItemCard{}, nil)
}

func CreateArgs() map[string]*graphql.ArgumentConfig {
	return types.GenerateGraphQlArgs(ItemCard{}, nil)
}
