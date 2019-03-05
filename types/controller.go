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
