// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package types

import (
	"github.com/graphql-go/graphql"
	"reflect"
)

func graphQlObject(name string, fields graphql.Fields) *graphql.Object {

	obj := graphql.NewObject(graphql.ObjectConfig{
		Name:   name,
		Fields: fields,
	})
	return obj
}

/*
GenerateGraphQlType generate a GraphQL type object from types defined in this application packages.

	Example:

		CustomerCard definition looks like this:

		'''
		type CustomerCard struct {
			No                          string `json:"No" required:"true"`
			Name                        string `json:"Name"`
			Address                     string `json:"Address"`
			...
		}
		'''

		// calling the generate function
		GenerateGraphQlType("CustomerCard", CustomerCard{}, nil)

		Returns:

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

		The returned object is a GraphQL object definition which will be utilized in the GrapQL
		main query or mutation

*/
func GenerateGraphQlType(name string, object interface{}, extraFields graphql.Fields) *graphql.Object {

	fields := graphql.Fields{}
	obj := reflect.TypeOf(object)

	for i := 0; i < obj.NumField(); i++ {
		_name := obj.Field(i).Tag.Get("json")
		_type := obj.Field(i).Type.String()

		if typeListContains(_type) {
			fields[_name] = attrMap[_type]
		}
	}

	if extraFields != nil {
		for key, value := range extraFields {
			fields[key] = value
		}
	}
	return graphQlObject(name, fields)
}

/*
GenerateGraphQlArgs creates GraphQl argument object definition

/*
CreateArgs function creates a GraphQl Object Type from the 'CustomerCard'

	Example:

		CustomerCard definition looks like this:

		'''
		type CustomerCard struct {
			No                          string `json:"No" required:"true"`
			Name                        string `json:"Name"`
			Address                     string `json:"Address"`
			...
		}
		'''

		// calling the generate function
		GenerateGraphQlArgs(CustomerCard{}, nil)

	Returns:

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

	example of required fields

		No	string `json:"No" required:"true"`

	and this will be translated to

		"No":	&graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},


	The returned GraphQl arguments will be used as a part of the main mutation
*/
func GenerateGraphQlArgs(object interface{}, extraFields map[string]*graphql.ArgumentConfig) map[string]*graphql.ArgumentConfig {

	args := map[string]*graphql.ArgumentConfig{}
	obj := reflect.TypeOf(object)
	for i := 0; i < obj.NumField(); i++ {

		readOnly := obj.Field(i).Tag.Get("readOnly")

		if readOnly == "true" {
			continue
		}

		name := obj.Field(i).Tag.Get("json")
		required := obj.Field(i).Tag.Get("required")
		_type := obj.Field(i).Type.String()

		if typeListContains(_type) {
			if required == "true" {
				args[name] = requiredArgMap[_type]
			} else {
				args[name] = argMap[_type]
			}
		}

	}

	if extraFields != nil {
		for key, value := range extraFields {
			args[key] = value
		}
	}
	return args
}
