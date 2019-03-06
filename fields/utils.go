// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package fields

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

/*
resolveFields gets required return fields from request query

	Example request query

	'''
	{
	  AssemblyBom {
		No
		Parent_Item_No
		Quantity_per
	  }
	}
	'''

	required fields are:

		"No",
		"Parent_Item_No",
		"Quantity_per"

Function will extract these fields from request and return a string slice as follows

	[No, Parent_Item_No, Quantity_per]

This will be used when creating a request for Microsoft Navision requesting these fields
to be retuned

Example request to Microsoft Navision:
	'''
	https://[ENDPOINT-BASE-URI]/Assembly_Bom?$select=No,Parent_Item_No,Quantity_per
	'''
*/
func resolveFields(params graphql.ResolveParams, selections []ast.Selection) ([]string, error) {
	var selected []string
	for _, s := range selections {
		switch t := s.(type) {
		case *ast.Field:
			selected = append(selected, s.(*ast.Field).Name.Value)
		case *ast.FragmentSpread:
			n := s.(*ast.FragmentSpread).Name.Value
			frag, ok := params.Info.Fragments[n]
			if !ok {
				return nil, fmt.Errorf("getSelectedFields: no fragment found with name %v", n)
			}
			sel, err := resolveFields(params, frag.GetSelectionSet().Selections)
			if err != nil {
				return nil, err
			}
			selected = append(selected, sel...)
		default:
			return nil, fmt.Errorf("getSelectedFields: found unexpected selection type %v", t)
		}
	}

	return selected, nil
}

/*
queryFields creates GraphQL Type fields for GraphQl query's

	Example:

	queryFields("assemblyBom", assemblybom.GetAll, assemblybom.Filter) would resolve to

		'''
		&graphql.Field{
			Type: graphql.NewList(types["assemblyBom"]), // types["assemblyBom"] --> assemblybom.createType()
			Args: filterArgs,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if len(p.Args) != 2 {
					return assemblybom.GetAll()
				}
				return assemblybom.Filter(p.Args)
			},
		}

*/
func queryFields(name string, getAll callback, filter callbackWithArgs) *graphql.Field {
	field := &graphql.Field{
		Type: graphql.NewList(types[name]),
		Args: filterArgs,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			fieldASTs := p.Info.FieldASTs

			if len(fieldASTs) == 0 {
				return nil, fmt.Errorf("getSelectedFields: ResolveParams has no fields")
			}
			fields, _ := resolveFields(p, fieldASTs[0].SelectionSet.Selections)

			if len(p.Args) != 2 {
				return getAll(fields)
			}
			return filter(fields, p.Args)
		},
	}
	return field
}

/*
createFields creates GraphQL Type fields for GraphQl mutation related to creating entities

	Example:

	createFields("customer", customer.Create) would resolve to

		'''
		&graphql.Field{
			Type: types["customer"], -- > customer.createType()
			Args: CustomerCardArgs,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				log.Printf("fetching Customer cards of company: %s", config.CompanyName)
				return customer.Create(p.Args)
			},
		}
		'''
*/
func createFields(name string, create callbackWithArgs) *graphql.Field {

	field := &graphql.Field{
		Type: types[name],
		Args: args[name],
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			fieldASTs := p.Info.FieldASTs

			if len(fieldASTs) == 0 {
				return nil, fmt.Errorf("getSelectedFields: ResolveParams has no fields")
			}
			fields, _ := resolveFields(p, fieldASTs[0].SelectionSet.Selections)

			return create(fields, p.Args)
		},
	}
	return field
}

/*
updateFields creates GraphQL Type fields for GraphQl mutation related to updating entities

	Example:

		updateFields("customer", customer.Update) would resolve to

		'''
		&graphql.Field{
			Type: types["customer"],  -- > customer.createType()
			Args: CustomerCardArgs,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				log.Printf("updating Customer cards of company: %s", config.CompanyName)
				return customer.Update(p.Args)
			},
		}
		'''
*/
func updateFields(name string, update callbackWithArgs) *graphql.Field {
	field := &graphql.Field{
		Type: types[name],
		Args: args[name],
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			fieldASTs := p.Info.FieldASTs

			if len(fieldASTs) == 0 {
				return nil, fmt.Errorf("getSelectedFields: ResolveParams has no fields")
			}
			fields, _ := resolveFields(p, fieldASTs[0].SelectionSet.Selections)

			return update(fields, p.Args)
		},
	}
	return field
}
