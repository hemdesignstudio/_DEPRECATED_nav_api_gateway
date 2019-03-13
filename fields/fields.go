package fields

import (
	"github.com/graphql-go/graphql"
)

type fieldType interface {
	GetCompany() string
	CreateType() *graphql.Object
	CreateArgs() map[string]*graphql.ArgumentConfig
	SetArgs(map[string]interface{})
	SetFields([]string)
	GetAll() (interface{}, error)
	Filter() (interface{}, error)
	Update() (interface{}, error)
	Create() (interface{}, error)
}

/*
 types contains a map of GraphQL Type objects of (assemblybom, Customer, item  .. etc)

Example GraphQL type for 'AssemblyBom':

	Example:
		'''
		graphql.NewObject(graphql.ObjectConfig{
				Name: "AssemblyBom",
				Fields: graphql.Fields{
					"Parent_Item_No":       &graphql.Field{Type: graphql.String},
					"No":                   &graphql.Field{Type: graphql.String},
					"Type":                 &graphql.Field{Type: graphql.String},
					...
				},
			})
		'''


*/

//filterArgs hold arguments used for filtering
var filterArgs = map[string]*graphql.ArgumentConfig{
	"key":   {Type: graphql.String},
	"value": {Type: graphql.String},
}

/*
args hold all create and update arguments for all mutation types

Example of GraphQl Argument Object for 'customer'

	Example:

		customer.CreateArgs() would resolve to the following

		'''
		map[string]*graphql.ArgumentConfig{
			"No":			&graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"Name":			&graphql.ArgumentConfig{Type: graphql.String},
			"Address":		&graphql.ArgumentConfig{Type: graphql.String},
			...
		}
		'''

	Hint:
		arguments are used to create or update entities,
		some arguments are required and hence in the CustomerCard type,
		tags can be noticed

		example of required fields

			No	string `json:"No" required:"true"`

		and this will be translated to

			"No":	&graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},


The returned GraphQl arguments will be used as a part of the main mutation
*/
