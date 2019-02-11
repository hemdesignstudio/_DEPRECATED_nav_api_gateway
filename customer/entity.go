package customer

import (
	"github.com/graphql-go/graphql"
	"github.com/nav-api-gateway/config"
	"github.com/nav-api-gateway/types"
)

var endpoint = config.CustomerCardWSEndpoint
var companyName = config.CompanyName

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

func CreateType() *graphql.Object {
	return types.GenerateGraphQlType("CustomerCard", CustomerCard{}, nil)
}
