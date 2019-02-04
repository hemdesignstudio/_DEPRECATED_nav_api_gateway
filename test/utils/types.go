package utils

type SliceOfMaps []map[string]interface{}
type MapOfSliceOfMaps map[string][]map[string]interface{}

type ErrorMessage struct {
	Message interface{} `json:"message"`
}

type QueryType struct {
	AssemblyBom  string
	CustomerCard string
	ItemCard     string
}

type MutationType struct {
	CreateCustomerCard string
	CreateItemCard     string

	UpdateCustomerCard string
	UpdateItemCard     string
}

var Query = QueryType{
	AssemblyBom:  "AssemblyBom",
	CustomerCard: "CustomerCard",
	ItemCard:     "ItemCard",
}

var Mutation = MutationType{
	CreateCustomerCard: "CreateCustomerCard",
	UpdateCustomerCard: "UpdateCustomerCard",

	CreateItemCard: "CreateItemCard",
	UpdateItemCard: "UpdateItemCard",
}
