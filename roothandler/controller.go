// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

/*
Package roothandler implements a simple package for HTTP Handler functions.


*/
package roothandler

import (
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
	"github.com/hem-nav-gateway/field"
	"log"
	"net/http"
)

// Handler creates a handler function for Graphql Schema
func Handler(company string) *gqlhandler.Handler {

	query := field.QueryType(company)
	mutation := field.MutationType(company)

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    query,
		Mutation: mutation,
	})

	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	handler := gqlhandler.New(&gqlhandler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: false,
	})

	return handler
}

// RootEndpoint is a handler function for the main endpoint
func RootEndpoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	company, err := pathVariables(vars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	handler := Handler(company)
	handler.ServeHTTP(w, r)
}
