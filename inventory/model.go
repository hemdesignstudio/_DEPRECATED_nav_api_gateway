package inventory

import (
	"github.com/graphql-go/graphql"
	"github.com/hem-nav-gateway/types"
)

/*


Example:
	'''
	<Soap:Envelope xmlns:Soap="http://schemas.xmlsoap.org/soap/envelope/">
		<Soap:Body>
			<GetInventory_Result xmlns="urn:microsoft-dynamics-schemas/codeunit/GetInventory">
				<retValues>
						<InventoryLine xmlns="urn:microsoft-dynamics-nav/xmlports/x50000">
						<ItemNo>10005</ItemNo>
						<Description>Hai Chair Mosaic Charcoal</Description>
						<ReceiptDate>2019-03-08</ReceiptDate>
						<AvailableDate>25.00</AvailableDate>
					</InventoryLine>
				</retValues>
			</GetInventory_Result>
		</Soap:Body>
	</Soap:Envelope>
	'''

*/
type Response struct {
	Body struct {
		GetInventoryResult struct {
			RetValues struct {
				InventoryLine []Model `xml:"InventoryLine"`
			} `xml:"retValues"`
		} `xml:"GetInventory_Result"`
	}
}

type Model struct {
	ItemNo        string `xml:"ItemNo" json:"ItemNo"`
	Description   string `xml:"Description" json:"Description"`
	ReceiptDate   string `xml:"ReceiptDate" json:"ReceiptDate"`
	AvailableDate string `xml:"AvailableDate" json:"AvailableDate"`
}

/*
CreateType function creates a GraphQl Object Type from the 'Inventory' type.

Example of GraphQl Object

	Example:
		'''
		graphql.NewObject(graphql.ObjectConfig{
				Name: "Inventory",
				Fields: graphql.Fields{
					"No":					&graphql.Field{Type: graphql.String},
					"Description":			&graphql.Field{Type: graphql.String},
					"Base_Unit_of_Measure":	&graphql.Field{Type: graphql.String},
					...
				},
			})
		'''

GraphQl Object is a map[string]*graphql.Field

The returned GraphQl Object Type will be used as a part of the main query
*/
func createType() *graphql.Object {
	return types.GenerateGraphQlType("Inventory", Model{}, nil)
}

/*
CreateArgs function creates a GraphQl Object Type from the 'Inventory'

Example of GraphQl Argument Object

	Example:
		'''
		map[string]*graphql.ArgumentConfig{
			"No":					&graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"Description":			&graphql.ArgumentConfig{Type: graphql.String},
			"Base_Unit_of_Measure":	&graphql.ArgumentConfig{Type: graphql.String},
			...
		}
		'''

	Hint:
		Arguments are used to create or update entities,
		some arguments are required and hence in the Inventory type,
		tags can be noticed

Example of required fields

	No	string `json:"No" required:"true"`

and this will be translated to

	"No":	&graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},


The returned GraphQl arguments will be used as a part of the main mutation
*/
func createArgs() map[string]*graphql.ArgumentConfig {
	return types.GenerateGraphQlArgs(Model{}, nil)
}
