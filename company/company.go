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
	Id          string `json:"Id"`
	Name        string `json:"Name"`
	DisplayName string `json:"DisplayName"`
}

func QueryType() *graphql.Object {
	companyFields := getCompanyFields()
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Company",
		Fields: graphql.Fields{
			"Id":                        companyFields["Id"],
			"Name":                      companyFields["Name"],
			"DisplayName":               companyFields["DisplayName"],
			"AssemblyBom":               getAssemblyBomFields(),
			"customerCard":              getCustomerCardFields(),
			"ItemCard":                  getItemCardFields(),
			"SalesOrder":                getSalesOrdersFields(),
			"SalesLine":                 getSalesLinesFields(),
			"PostedSalesShipment":       getPostShipFields(),
			"SalesInvoice":              getSalesInvoiceFields(),
			"UpdateCustomerCard":        updateCustomerCardFields(),
			"UpdateItemCard":            updateItemCardFields(),
			"UpdateSalesOrder":          updateSalesOrderFields(),
			"UpdateSalesLine":           updateSalesLineFields(),
			"UpdatePostedSalesShipment": updatePostShipFields(),
			"UpdateSalesInvoice":        updateSalesInvoiceFields(),
		},
	})
}

func GetCompanyByName() (*Company, error) {
	url := config.BaseUrl + config.CompanyEndpoint + fmt.Sprintf("('%s')", config.CompanyName)
	resultByte, err := request.GET(url)
	response := Company{}
	err = json.Unmarshal(resultByte, &response)
	if err != nil {
		return nil, errors.New("could not unmarshal data")
	}
	return &response, err
}
