package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
	"github.com/nav-api-gateway/company"
	"github.com/nav-api-gateway/config"
	"log"
	"net/http"
)

func main() {
	companyQuery := company.QueryType()
	query := graphql.NewObject(createQueryType(companyQuery))

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: query,
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

	http.Handle("/graphql", handler)
	fmt.Println("Server started at http://localhost:6789/graphql")
	log.Fatal(http.ListenAndServe(config.Host, nil))
}

func createQueryType(companyType *graphql.Object) graphql.ObjectConfig {
	rootQuery := graphql.ObjectConfig{
		Name: "QueryType",
		Fields: graphql.Fields{
			"company": &graphql.Field{
				Type: companyType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					log.Printf("fetching companies with name: %s", config.CompanyName)
					return company.GetCompanyByName()
				},
			},
		}}
	return rootQuery
}
