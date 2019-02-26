package utils

import (
	"github.com/graphql-go/graphql"
)

type Args map[string]*graphql.ArgumentConfig
type getAllFunction func() (interface{}, error)
type filterFunction func(map[string]interface{}) (interface{}, error)
type createFunction func(map[string]interface{}) (interface{}, error)
type UpdateFunction func(map[string]interface{}) (interface{}, error)

func queryFields(name string, getAll getAllFunction, filter filterFunction) *graphql.Field {
	field := &graphql.Field{
		Type: graphql.NewList(types[name]),
		Args: filterArgs,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if len(p.Args) != 2 {
				return getAll()
			}
			return filter(p.Args)
		},
	}
	return field
}

func createFields(name string, args map[string]*graphql.ArgumentConfig, create createFunction) *graphql.Field {
	field := &graphql.Field{
		Type: types[name],
		Args: args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return create(p.Args)
		},
	}
	return field
}

func updateFields(name string, args map[string]*graphql.ArgumentConfig, update UpdateFunction) *graphql.Field {
	field := &graphql.Field{
		Type: types[name],
		Args: args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return update(p.Args)
		},
	}
	return field
}
