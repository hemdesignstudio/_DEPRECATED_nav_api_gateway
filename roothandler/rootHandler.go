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
	"github.com/hem-nav-gateway/errorhandler"
	"github.com/hem-nav-gateway/fields"
	"log"
	"net/http"
)

var companies = []string{"test", "us", "europe"}

func checkCompany(companyName string) bool {
	for _, elem := range companies {
		if elem == companyName {
			return true
		}
	}
	return false
}

// pathVariables extracts vars from URL path
func pathVariables(vars map[string]string) (string, error) {
	company := vars["company"]

	if checkCompany(company) {
		return company, nil
	}
	return "", errorhandler.CompanyDoesNotExist(company)
}

// Handler creates a handler function for Graphql Schema
func Handler(company string) *gqlhandler.Handler {

	query := fields.QueryType(company)
	mutation := fields.MutationType(company)

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
