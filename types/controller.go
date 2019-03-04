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
		elemName := obj.Field(i).Tag.Get("json")
		elemType := obj.Field(i).Type.String()

		if typeListContains(elemType) {
			fields[elemName] = attrMap[elemType]
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
	t := reflect.TypeOf(object)
	for i := 0; i < t.NumField(); i++ {
		elemName := t.Field(i).Tag.Get("json")
		elemRequired := t.Field(i).Tag.Get("required")
		elemType := t.Field(i).Type.String()

		if typeListContains(elemType) {
			if elemRequired == "true" {
				args[elemName] = requiredArgMap[elemType]
			} else {
				args[elemName] = argMap[elemType]
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
