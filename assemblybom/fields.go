package assemblybom

import "github.com/graphql-go/graphql"

func CreateType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "AssemblyBom",
		Fields: graphql.Fields{
			"Parent_Item_No":       &graphql.Field{Type: graphql.String},
			"No":                   &graphql.Field{Type: graphql.String},
			"Type":                 &graphql.Field{Type: graphql.String},
			"Quantity_per":         &graphql.Field{Type: graphql.Float},
			"Unit_of_Measure_Code": &graphql.Field{Type: graphql.String},
		},
	})
}
