package salesline

import "github.com/graphql-go/graphql"

func CreateType(name string) *graphql.Object {

	fields := graphql.Fields{
		"No":                    &graphql.Field{Type: graphql.String},
		"Document_No":           &graphql.Field{Type: graphql.String},
		"Type":                  &graphql.Field{Type: graphql.String},
		"Description":           &graphql.Field{Type: graphql.String},
		"Reserve":               &graphql.Field{Type: graphql.String},
		"Quantity":              &graphql.Field{Type: graphql.Int},
		"Reserved_Quantity":     &graphql.Field{Type: graphql.Int},
		"Unit_of_Measure_Code":  &graphql.Field{Type: graphql.String},
		"Unit_Price":            &graphql.Field{Type: graphql.Float},
		"Line_Amount":           &graphql.Field{Type: graphql.Float},
		"Line_Discount_Percent": &graphql.Field{Type: graphql.Float},
		"Line_Discount_Amount":  &graphql.Field{Type: graphql.Float},
		"Prepayment_Percent":    &graphql.Field{Type: graphql.Float},
		"Prepmt_Line_Amount":    &graphql.Field{Type: graphql.Float},
		"Qty_to_Ship":           &graphql.Field{Type: graphql.Int},
		"Quantity_Shipped":      &graphql.Field{Type: graphql.Int},
		"Qty_to_Invoice":        &graphql.Field{Type: graphql.Int},
		"Quantity_Invoiced":     &graphql.Field{Type: graphql.Int},
	}

	return graphql.NewObject(graphql.ObjectConfig{
		Name:   name,
		Fields: fields,
	})
}
