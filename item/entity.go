// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

/*
Package item implements a simple package for handling all graphql
operations related to Microsoft Navision ItemCard page.

Package has a type "ItemCard" where all the fields related to ItemCard are defined.

	type ItemCard struct {
		No                        string  `json:"No" required:"true"`
		Description               string  `json:"Description"`
		BaseUnitOfMeasure         string  `json:"Base_Unit_of_Measure"`
		...
	}



GraphQl Object Type along with its fields, arguments and attributes are generated
from the Item type when "CreateType" method is invoked.
*/
package item

import (
	"github.com/graphql-go/graphql"
	"github.com/hem-nav-gateway/config"
	"github.com/hem-nav-gateway/types"
)

//Microsoft Navision endpoint path for ItemCard page
var endpoint = config.ItemCardEndpoint

/*
Response is utilized as Microsoft Navision returns a list of objects
when requesting ItemCard, It is utilized for JSON decoding

example response from Navision

	{
		"value": [
			{
				"No": "1234"
				"Description": "this object is great",
				"Base_Unit_of_Measure": "ABM",
				...
			},
			{
				"No": "2345"
				"Description": "this object is amazing",
				"Base_Unit_of_Measure": "CMA",
				...
			},
			{
			...

			},
	}

*/
type Response struct {
	Value []ItemCard `json:"value"`
}

type ItemCard struct {
	No                        string  `json:"No" required:"true"`
	Description               string  `json:"Description"`
	BaseUnitOfMeasure         string  `json:"Base_Unit_of_Measure"`
	AssemblyBOM               bool    `json:"Assembly_BOM"`
	ItemCategoryCode          string  `json:"Item_Category_Code"`
	ProductGroupCode          string  `json:"Product_Group_Code"`
	Inventory                 int     `json:"Inventory"`
	QtyOnPurchOrder           int     `json:"Qty_on_Purch_Order"`
	QtyOnSalesOrder           int     `json:"Qty_on_Sales_Order"`
	LastDateModified          string  `json:"Last_Date_Modified"`
	FreightType               string  `json:"Freight_Type"`
	AssemblyPolicy            string  `json:"Assembly_Policy"`
	CountryRegionOfOriginCode string  `json:"Country_Region_of_Origin_Code"`
	NetWeight                 float32 `json:"Net_Weight"`
	GrossWeight               float32 `json:"Gross_Weight"`
	UnitVolume                float32 `json:"Unit_Volume"`
	Length                    float32 `json:"Length"`
	Width                     float32 `json:"Width"`
	Height                    float32 `json:"Height"`
	Designer                  string  `json:"Designer"`
	WebStatus                 string  `json:"Web_Status"`
}

/*
CreateType function creates a GraphQl Object Type from the 'ItemCard' type.

example of GraphQl Object

	graphql.NewObject(graphql.ObjectConfig{
			Name: "ItemCard",
			Fields: graphql.Fields{
				"No":					&graphql.Field{Type: graphql.String},
				"Description":			&graphql.Field{Type: graphql.String},
				"Base_Unit_of_Measure":	&graphql.Field{Type: graphql.String},
				...
			},
		})

GraphQl Object is a map[string]*graphql.Field

The returned GraphQl Object Type will be used as a part of the main query
*/
func CreateType() *graphql.Object {
	return types.GenerateGraphQlType("ItemCard", ItemCard{}, nil)
}

/*
CreateArgs function creates a GraphQl Object Type from the 'ItemCard'

example of GraphQl Argument Object

	map[string]*graphql.ArgumentConfig{
		"No":					&graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"Description":			&graphql.ArgumentConfig{Type: graphql.String},
		"Base_Unit_of_Measure":	&graphql.ArgumentConfig{Type: graphql.String},
		...
	}

Hint: arguments are used to create or update entities,
some arguments are required and hence in the ItemCard type,
tags can be noticed

example of required fields

	No	string `json:"No" required:"true"`

and this will be translated to

	"No":	&graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},


The returned GraphQl Object Type will be used as a part of the main query
*/
func CreateArgs() map[string]*graphql.ArgumentConfig {
	return types.GenerateGraphQlArgs(ItemCard{}, nil)
}
