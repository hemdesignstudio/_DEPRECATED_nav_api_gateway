package fields

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/hem-nav-gateway/assemblybom"
	"github.com/hem-nav-gateway/customer"
	"github.com/hem-nav-gateway/item"
	"github.com/hem-nav-gateway/postship"
	"github.com/hem-nav-gateway/salesinvoice"
	"github.com/hem-nav-gateway/salesline"
	"github.com/hem-nav-gateway/salesorder"
)

type callback func(interface{}) (interface{}, error)
type callbackWithArgs func(interface{}, interface{}) (interface{}, error)
type Args map[string]map[string]*graphql.ArgumentConfig

var types = map[string]*graphql.Object{
	"assemblyBom":  assemblybom.CreateType(),
	"customer":     customer.CreateType(),
	"item":         item.CreateType(),
	"salesOrder":   salesorder.CreateType(),
	"salesLine":    salesline.CreateType("SalesLine"),
	"postShip":     postship.CreateType(),
	"salesInvoice": salesinvoice.CreateType(),
}

var filterArgs = map[string]*graphql.ArgumentConfig{
	"key":   {Type: graphql.String},
	"value": {Type: graphql.String},
}

var args = Args{
	"customer":     customer.CreateArgs(),
	"item":         item.CreateArgs(),
	"salesOrder":   salesorder.CreateArgs(),
	"salesLine":    salesline.CreateArgs(),
	"postShip":     postship.CreateArgs(),
	"salesInvoice": salesinvoice.CreateArgs(),
}

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

func QueryType() *graphql.Object {
	query := graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"AssemblyBom":         queryFields("assemblyBom", assemblybom.GetAll, assemblybom.Filter),
			"CustomerCard":        queryFields("customer", customer.GetAll, customer.Filter),
			"ItemCard":            queryFields("item", item.GetAll, item.Filter),
			"SalesOrder":          queryFields("salesOrder", salesorder.GetAll, salesorder.Filter),
			"SalesLine":           queryFields("salesLine", salesline.GetAll, salesline.Filter),
			"PostedSalesShipment": queryFields("postShip", postship.GetAll, postship.Filter),
			"SalesInvoice":        queryFields("salesInvoice", salesinvoice.GetAll, salesinvoice.Filter),
		},
	}
	return graphql.NewObject(query)
}

func MutationType() *graphql.Object {
	mutation := graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"CreateCustomerCard": createFields("customer", customer.Create),
			"CreateItemCard":     createFields("item", item.Create),
			"CreateSalesOrder":   createFields("salesOrder", salesorder.Create),
			"CreateSalesLine":    createFields("salesLine", salesline.Create),
			"CreateSalesInvoice": createFields("salesInvoice", salesinvoice.Create),
			"UpdateCustomerCard": updateFields("customer", customer.Update),
			"UpdateItemCard":     updateFields("item", item.Update),
			"UpdateSalesOrder":   updateFields("salesOrder", salesorder.Update),
			"UpdateSalesLine":    updateFields("salesLine", salesline.Update),
			"UpdateSalesInvoice": updateFields("salesInvoice", salesinvoice.Update),
		},
	}
	return graphql.NewObject(mutation)
}