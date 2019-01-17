package assemblybom

import (
	"github.com/nav-api-gateway/config"
)

var endpoint = config.AssemblyBomEndpoint
var companyName = config.CompanyName

type Response struct {
	Value []AssemblyBom `json:"value"`
}

type AssemblyBom struct {
	ParentItemNo      string  `json:"Parent_Item_No"`
	No                string  `json:"No"`
	Type              string  `json:"Type"`
	QuantityPer       float64 `json:"Quantity_per"`
	UnitOfMeasureCode string  `json:"Unit_of_Measure_Code"`
}
