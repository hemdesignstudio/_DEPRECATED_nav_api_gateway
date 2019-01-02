package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
	"log"
	"net/http"
	"projects/graphql/assemblybom"
	"projects/graphql/company"
	"projects/graphql/config"
	"projects/graphql/customer"
)

func main() {
	conf := config.GetConfig()
	// our schema is generated here
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(
			createQueryType(
				company.CreateCompanyType(
					customer.CreateCustomerCardType(),
					assemblybom.CreateAssemblyBomType(),
				),
			),
		),
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
	fmt.Println("Server started at http://localhost:3000/graphql")
	log.Fatal(http.ListenAndServe(conf.Host, nil))
}

func createQueryType(companyType *graphql.Object) graphql.ObjectConfig {
	return graphql.ObjectConfig{Name: "QueryType", Fields: graphql.Fields{
		"company": &graphql.Field{
			Type: companyType,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				name := p.Args["name"]
				s, _ := name.(string)
				log.Printf("fetching companies with name: %s", s)
				return company.GetCompanyByName(s)
			},
		},
	}}
}
