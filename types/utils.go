// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package types

import "github.com/graphql-go/graphql"

var elementTypes = []string{

	"string",
	"float32",
	"float64",
	"int",
	"bool",
}

var argMap = map[string]*graphql.ArgumentConfig{
	"string":  arg.STRING,
	"float32": arg.FLOAT,
	"float64": arg.FLOAT,
	"int":     arg.INT,
	"bool":    arg.BOOL,
}

var requiredArgMap = map[string]*graphql.ArgumentConfig{
	"string":  requiredArg.STRING,
	"float32": requiredArg.FLOAT,
	"float64": requiredArg.FLOAT,
	"int":     requiredArg.INT,
	"bool":    requiredArg.BOOL,
}

var fieldMap = map[string]*graphql.Field{
	"string":  field.STRING,
	"float32": field.FLOAT,
	"float64": field.FLOAT,
	"int":     field.INT,
	"bool":    field.BOOL,
}

// typeListContains checks if element type exits in
// the accepted element types
func typeListContains(elementType string) bool {

	for _, elem := range elementTypes {
		if elementType == elem {
			return true
		}
	}
	return false
}
