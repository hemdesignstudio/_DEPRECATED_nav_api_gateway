// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

/*
Package types implements a package generating GraphQL Types and Mutation Arguments.


*/
package types

import "github.com/graphql-go/graphql"

type attrType struct {
	STRING *graphql.Field
	INT    *graphql.Field
	FLOAT  *graphql.Field
	BOOL   *graphql.Field
}

type argType struct {
	STRING *graphql.ArgumentConfig
	INT    *graphql.ArgumentConfig
	FLOAT  *graphql.ArgumentConfig
	BOOL   *graphql.ArgumentConfig
}

var attr = attrType{
	STRING: &graphql.Field{Type: graphql.String},
	INT:    &graphql.Field{Type: graphql.Int},
	FLOAT:  &graphql.Field{Type: graphql.Float},
	BOOL:   &graphql.Field{Type: graphql.Boolean},
}

var arg = argType{
	STRING: &graphql.ArgumentConfig{Type: graphql.String},
	INT:    &graphql.ArgumentConfig{Type: graphql.Int},
	FLOAT:  &graphql.ArgumentConfig{Type: graphql.Float},
	BOOL:   &graphql.ArgumentConfig{Type: graphql.Boolean},
}

var requiredArg = argType{
	STRING: &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
	INT:    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
	FLOAT:  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Float)},
	BOOL:   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Boolean)},
}
