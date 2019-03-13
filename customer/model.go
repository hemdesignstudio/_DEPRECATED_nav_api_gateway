// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

/*
Package customer implements a simple package for handling all graphql
operations related to Microsoft Navision CustomerCardWS page.

Package has a type "CustomerCard" where all the fields related to CustomerCardWS are defined.

	Example:
		'''
		type CustomerCard struct {
			No                          string `json:"No" required:"true"`
			Name                        string `json:"Name"`
			Address                     string `json:"Address"`
			Address2                    string `json:"Address_2"`
			...
		}
		'''

GraphQl Object Type along with its fields, arguments and attributes are generated
from the CustomerCard type when "CreateType" method is invoked.
*/
package customer

import (
	"github.com/graphql-go/graphql"
	"github.com/hem-nav-gateway/config"
	"github.com/hem-nav-gateway/types"
)

// Microsoft Navision endpoint path for CustomerCard page
var endpoint = config.CustomerCardWSEndpoint

/*
Response is utilized as Microsoft Navision returns a list of objects
when requesting CustomerCardWS, It is utilized for JSON decoding

Example response from Navision

	Example:
		'''
		{
			"value": [
				{
					"No": "1234"
					"Name": "John",
					"Address": "some address",
					"Address_2": "yet another address",
					...
				},
				{
					"No": "2345"
					"Name": "Doe",
					"Address": "some address",
					"Address_2": "yet another address",
					...
				},
				{
				...

				},
		}
		'''

*/
type Response struct {
	Value []Model `json:"value"`
}

type Model struct {
	No                          string `json:"No" required:"true"`
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

/*
CreateType function creates a GraphQl Object Type from the 'CustomerCard' type.

Example of GraphQl Object

	Example:
		'''
		graphql.NewObject(graphql.ObjectConfig{
				Name: "CustomerCard",
				Fields: graphql.Fields{
					"No":				&graphql.Field{Type: graphql.String},
					"Name":				&graphql.Field{Type: graphql.String},
					"Address":			&graphql.Field{Type: graphql.String},
					...
				},
			})
		'''

GraphQl Object is a map[string]*graphql.Field

The returned GraphQl Object Type will be used as a part of the main query
*/
func createType() *graphql.Object {

	return types.GenerateGraphQlType("CustomerCard", Model{}, nil)
}

/*
CreateArgs function creates a GraphQl Object Type from the 'CustomerCard'

Example of GraphQl Argument Object

	Example:
		'''
		map[string]*graphql.ArgumentConfig{
			"No":			&graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"Name":			&graphql.ArgumentConfig{Type: graphql.String},
			"Address":		&graphql.ArgumentConfig{Type: graphql.String},
			...
		}
		'''

Hint: arguments are used to create or update entities,
some arguments are required and hence in the CustomerCard type,
tags can be noticed

Example of required fields

	No	string `json:"No" required:"true"`

and this will be translated to

	"No":	&graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},


The returned GraphQl arguments will be used as a part of the main mutation
*/
func createArgs() map[string]*graphql.ArgumentConfig {

	return types.GenerateGraphQlArgs(Model{}, nil)
}
