package customer

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/nav-api-gateway/config"
	"github.com/nav-api-gateway/request"
)

type Response struct {
	Value []CustomerCard `json:"value"`
}

type CustomerCard struct {
	No                          string `json:"No"`
	Name                        string `json:"Name"`
	Address                     string `json:"Address"`
	Address2                    string `json:"Address_2"`
	PostCode                    string `json:"Post_Code"`
	City                        string `json:"City"`
	CountryRegionCode           string `json:"Country_Region_Code"`
	PhoneNo                     string `json:"Phone_No"`
	Contact                     string `json:"Contact"`
	VatRegistrationNumber       string `json:"VAT_Registration_No"`
	CustomerPostingGroup        string `json:"Customer_Posting_Group"`
	GeneralBusinessPostingGroup string `json:"Gen_Bus_Posting_Group"`
	VatBusinessPostingGroup     string `json:"VAT_Bus_Posting_Group"`
	SalesTypeCode               string `json:"Global_Dimension_2_Code"`
	CustomerPriceGroup          string `json:"Customer_Price_Group"`
	CustomerDiscountGroup       string `json:"Customer_Disc_Group"`
	PaymentTermsCode            string `json:"Payment_Terms_Code"`
	CurrencyCode                string `json:"Currency_Code"`
	LanguageCode                string `json:"Language_Code"`
	WebEmail                    string `json:"Web_E_Mail"`
	WebEnabled                  bool   `json:"Web_Customer"`
}

func CreateCustomerCardType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "CustomerCard",
		Fields: graphql.Fields{
			"No":                      &graphql.Field{Type: graphql.String},
			"Name":                    &graphql.Field{Type: graphql.String},
			"Address":                 &graphql.Field{Type: graphql.String},
			"Address_2":               &graphql.Field{Type: graphql.String},
			"Post_Code":               &graphql.Field{Type: graphql.String},
			"City":                    &graphql.Field{Type: graphql.String},
			"Country_Region_Code":     &graphql.Field{Type: graphql.String},
			"Phone_No":                &graphql.Field{Type: graphql.String},
			"Contact":                 &graphql.Field{Type: graphql.String},
			"VAT_Registration_No":     &graphql.Field{Type: graphql.String},
			"Customer_Posting_Group":  &graphql.Field{Type: graphql.String},
			"Gen_Bus_Posting_Group":   &graphql.Field{Type: graphql.String},
			"VAT_Bus_Posting_Group":   &graphql.Field{Type: graphql.String},
			"Global_Dimension_2_Code": &graphql.Field{Type: graphql.String},
			"Customer_Price_Group":    &graphql.Field{Type: graphql.String},
			"Customer_Disc_Group":     &graphql.Field{Type: graphql.String},
			"Payment_Terms_Code":      &graphql.Field{Type: graphql.String},
			"Currency_Code":           &graphql.Field{Type: graphql.String},
			"Language_Code":           &graphql.Field{Type: graphql.String},
			"Web_E_Mail":              &graphql.Field{Type: graphql.String},
			"Web_Customer":            &graphql.Field{Type: graphql.Boolean},
		},
	})
}

func GetCustomerCardByCompanyName(name string) ([]CustomerCard, error) {
	conf := config.GetConfig()
	url := conf.BaseUrl + conf.CompanyEndpoint + fmt.Sprintf("('%s')"+conf.CustomerCardWSEndpoint, name)
	resultByte, err := request.GET(url)
	res := Response{}
	err = json.Unmarshal(resultByte, &res)

	if err != nil {
		return nil, errors.New("could not unmarshal data")
	}
	return res.Value, nil
}

func GetCustomerCardByNo(companyName string, no string) (*CustomerCard, error) {
	conf := config.GetConfig()
	url := conf.BaseUrl +
		conf.CompanyEndpoint + fmt.Sprintf("('%s')", companyName) +
		conf.CustomerCardWSEndpoint + fmt.Sprintf("('%s')", no)

	resultByte, err := request.GET(url)
	response := CustomerCard{}
	err = json.Unmarshal(resultByte, &response)

	if err != nil {
		return nil, errors.New("could not unmarshal data")
	}
	return &response, nil
}
