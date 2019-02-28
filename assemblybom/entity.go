package assemblybom

import (
	"github.com/graphql-go/graphql"
	"github.com/hem-nav-gateway/config"
	"github.com/hem-nav-gateway/types"
)

var endpoint = config.AssemblyBomEndpoint

type (
	Response struct {
		Value []AssemblyBom `json:"value"`
	}

	AssemblyBom struct {
		No                string  `json:"No" required:"true"`
		ParentItemNo      string  `json:"Parent_Item_No"`
		Type              string  `json:"Type"`
		QuantityPer       float64 `json:"Quantity_per"`
		UnitOfMeasureCode string  `json:"Unit_of_Measure_Code"`
	}
)

func CreateType() *graphql.Object {
	assemblyBom := &AssemblyBom{}
	return types.GenerateGraphQlType("AssemblyBom", *assemblyBom, nil)
}
