package assemblybom

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/nav-api-gateway/config"
	"github.com/nav-api-gateway/request"
)

var conf = config.GetConfig()

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

func CreateAssemblyBomType() *graphql.Object {
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

func GetAssemblyBomByCompanyName(name string) ([]AssemblyBom, error) {
	url := conf.BaseUrl + conf.CompanyEndpoint + fmt.Sprintf("('%s')"+conf.AssemblyBomEndpoint, name)
	resultByte, err := request.GET(url)
	res := Response{}
	err = json.Unmarshal(resultByte, &res)

	if err != nil {
		return nil, errors.New("could not unmarshal data")
	}
	return res.Value, nil
}

func GetAssemblyBomByFilter(companyName string, args map[string]interface{}) ([]AssemblyBom, error) {
	resByte := request.Filter(companyName, conf.AssemblyBomEndpoint, args)
	res := Response{}
	err := json.Unmarshal(resByte, &res)
	if err != nil {
		return nil, errors.New("could not unmarshal data")
	}
	return res.Value, nil
}
