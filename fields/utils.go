package fields

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

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
