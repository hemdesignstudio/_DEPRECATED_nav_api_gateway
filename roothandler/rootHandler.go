// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

/*
Package roothandler implements a simple package for HTTP Handler functions.


*/
package roothandler

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
	"github.com/hem-nav-gateway/config"
	"github.com/hem-nav-gateway/errorhandler"
	"github.com/hem-nav-gateway/fields"
	"log"
	"net/http"
)

// pathVariables extracts vars from URL path
func pathVariables(vars map[string]string) (string, error) {
	companyPath := vars["company"]
	if vars["company"] == "test" {
		return config.TestCompanyName, nil
	}

	return "", errorhandler.CompanyDoesNotExist(companyPath)
}

// Handler creates a handler function for Graphql Schema
func Handler() *gqlhandler.Handler {

	query := fields.QueryType()
	mutation := fields.MutationType()

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
	companyName, err := pathVariables(vars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Println(companyName)
	handler := Handler()
	handler.ServeHTTP(w, r)
}
