package assemblybom

import (
	"github.com/graphql-go/graphql"
	"github.com/hem-nav-gateway/types"
)

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

func CreateType() *graphql.Object {
	return types.GenerateGraphQlType("AssemblyBom", AssemblyBom{}, nil)
}
