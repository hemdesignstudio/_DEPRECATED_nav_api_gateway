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
	companyType := company.CreateCompanyType()
	query := graphql.NewObject(createQueryType(companyType))
	mutation := graphql.NewObject(createMutationType(companyType))

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
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},

				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					name := p.Args["name"]
					config.CompanyName, _ = name.(string)
					log.Printf("fetching companies with name: %s", config.CompanyName)
					return company.GetCompanyByName()
				},
			},
		}}
	return rootQuery
}

func createMutationType(companyType *graphql.Object) graphql.ObjectConfig {
	rootMutation := graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"company": &graphql.Field{
				Type: companyType, // the return type for this field
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// marshall and cast the argument value
					name := p.Args["name"]
					config.CompanyName, _ = name.(string)
					return company.GetCompanyByName()
				},
			},
		},
	}
	return rootMutation
}
