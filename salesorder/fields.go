package salesorder

import (
	"github.com/graphql-go/graphql"
	//"github.com/nav-api-gateway/company"
	"github.com/nav-api-gateway/salesline"
)

var types = map[string]*graphql.Object{
	"salesLine": salesline.CreateSalesLineType(),
}

/*
func getSalesLinesFields() *graphql.Field {
	field := &graphql.Field{
		Type: graphql.NewList(types["salesLine"]),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			company, _ := p.Source.(*company.Company)
			log.Printf("fetching sales lines of company: %s", company.Name)
			if len(p.Args) != 2 {
				return salesline.GetSalesLineByCompanyName(company.Name)
			}

			return salesline.GetSalesLineByFilter(company.Name, p.Args)
		},
	}
	return field
}
*/
