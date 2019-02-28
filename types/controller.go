package types

import (
	"github.com/graphql-go/graphql"
	"reflect"
	"strings"
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
	t := reflect.TypeOf(object)

	for i := 0; i < t.NumField(); i++ {
		elemName := t.Field(i).Tag.Get("json")
		elemType := t.Field(i).Type.String()

		if elemType == "string" {
			fields[elemName] = attr.STRING

		} else if strings.Contains(elemType, "float") {
			fields[elemName] = attr.FLOAT

		} else if elemType == "int" {
			fields[elemName] = attr.INT

		} else if elemType == "bool" {
			fields[elemName] = attr.BOOL

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

		if elemRequired == "true" {
			if elemType == "string" {
				args[elemName] = requiredArg.STRING

			} else if strings.Contains(elemType, "float") {
				args[elemName] = requiredArg.FLOAT

			} else if elemType == "int" {
				args[elemName] = requiredArg.INT

			} else if elemType == "bool" {
				args[elemName] = requiredArg.BOOL
			}
		} else {
			if elemType == "string" {
				args[elemName] = arg.STRING

			} else if strings.Contains(elemType, "float") {
				args[elemName] = arg.FLOAT

			} else if elemType == "int" {
				args[elemName] = arg.INT

			} else if elemType == "bool" {
				args[elemName] = arg.BOOL
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
