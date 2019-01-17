package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
	"github.com/nav-api-gateway/config"
	"github.com/nav-api-gateway/utils"
	"log"
	"net/http"
)

func main() {
	query := utils.QueryType()
	mutation := utils.MutationType()

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

	http.Handle(config.Endpoint+config.Version, handler)
	fmt.Println("Server started at http://localhost:6789/graphql/v0.1.0")
	log.Fatal(http.ListenAndServe(config.Host, nil))
}
