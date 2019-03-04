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

var attrMap = map[string]*graphql.Field{
	"string":  attr.STRING,
	"float32": attr.FLOAT,
	"float64": attr.FLOAT,
	"int":     attr.INT,
	"bool":    attr.BOOL,
}

func typeListContains(elementType string) bool {

	for _, elem := range elementTypes {
		if elementType == elem {
			return true
		}
	}
	return false
}
