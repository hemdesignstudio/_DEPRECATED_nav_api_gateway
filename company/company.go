package company

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/nav-api-gateway/config"
	"github.com/nav-api-gateway/request"
)

type Company struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

func CreateCompanyType() *graphql.Object {

	companyFields := getCompanyFields()

	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Company",
		Fields: graphql.Fields{

			"id":          companyFields["id"],
			"name":        companyFields["name"],
			"displayName": companyFields["displayName"],

			"AssemblyBom":         getAssemblyBomFields(),
			"customerCard":        getCustomerCardFields(),
			"ItemCard":            getItemCardFields(),
			"SalesOrder":          getSalesOrdersFields(),
			"SalesLine":           getSalesLinesFields(),
			"PostedSalesShipment": getPostShipFields(),
		},
	})
}

func GetCompanyByName(name string) (*Company, error) {
	url := config.BaseUrl + config.CompanyEndpoint + fmt.Sprintf("('%s')", name)
	resultByte, err := request.GET(url)
	response := Company{}

	err = json.Unmarshal(resultByte, &response)
	if err != nil {
		return nil, errors.New("could not unmarshal data")
	}
	return &response, err
}
