package utils

type QueryType struct {
	AssemblyBom  string
	CustomerCard string
}

type MutationType struct {
	CreateCustomerCard string
}

var Query = QueryType{
	AssemblyBom:  "AssemblyBom",
	CustomerCard: "CustomerCard",
}

var Mutation = MutationType{
	CreateCustomerCard: "CreateCustomerCard",
}
