package assemblybom

import (
	"github.com/graphql-go/graphql"
	"github.com/nav-api-gateway/config"
	"github.com/nav-api-gateway/request"
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

func CreateType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "AssemblyBom",
		Fields: graphql.Fields{
			"Parent_Item_No":       &graphql.Field{Type: graphql.String},
			"No":                   &graphql.Field{Type: graphql.String},
			"Type":                 &graphql.Field{Type: graphql.String},
			"Quantity_per":         &graphql.Field{Type: graphql.Float},
			"Unit_of_Measure_Code": &graphql.Field{Type: graphql.String},
		},
	})
}

func GetAll() (interface{}, error) {
	res := Response{}
	return request.GetAllByCompanyName(companyName, endpoint, res)
}

func Filter(args map[string]interface{}) (interface{}, error) {
	res := Response{}
	return request.FilterByArgs(companyName, endpoint, args, res)
}
