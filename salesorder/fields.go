package salesorder

import (
	"github.com/graphql-go/graphql"
	"github.com/nav-api-gateway/config"
	"github.com/nav-api-gateway/salesline"
	"log"
)

var types = map[string]*graphql.Object{
	"salesLine": salesline.CreateSalesLineType(),
}

func getSalesLinesFields() *graphql.Field {
	field := &graphql.Field{
		Type: graphql.NewList(types["salesLine"]),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			log.Printf("fetching sales lines of company: %s", config.CompanyName)
			salesOrder, _ := p.Source.(SalesOrder)
			print(salesOrder.No)
			p.Args["key"] = "Document_No"
			p.Args["value"] = salesOrder.No
			return salesline.Filter(p.Args)
		},
	}
	return field
}
