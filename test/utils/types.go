package utils

type SliceOfMaps []map[string]interface{}
type MapOfSliceOfMaps map[string][]map[string]interface{}

type QueryType struct {
	AssemblyBom  string
	CustomerCard string
}

type MutationType struct {
	CreateCustomerCard string
	UpdateCustomerCard string
}

var Query = QueryType{
	AssemblyBom:  "AssemblyBom",
	CustomerCard: "CustomerCard",
}

var Mutation = MutationType{
	CreateCustomerCard: "CreateCustomerCard",
	UpdateCustomerCard: "UpdateCustomerCard",
}
