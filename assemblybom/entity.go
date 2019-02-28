package assemblybom

import (
	"github.com/graphql-go/graphql"
	"github.com/hem-nav-gateway/config"
	"github.com/hem-nav-gateway/types"
)

var endpoint = config.AssemblyBomEndpoint

type response struct {
	Value []AssemblyBom `json:"value"`
}

type AssemblyBom struct {
	No                string  `json:"No"`
	ParentItemNo      string  `json:"Parent_Item_No"`
	Type              string  `json:"Type"`
	QuantityPer       float64 `json:"Quantity_per"`
	UnitOfMeasureCode string  `json:"Unit_of_Measure_Code"`
}

func CreateType() *graphql.Object {
	return types.GenerateGraphQlType("AssemblyBom", AssemblyBom{}, nil)
}
