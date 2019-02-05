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
	SalesOrder   string
	SalesInvoice string
}

type MutationType struct {
	CreateCustomerCard string
	CreateItemCard     string
	CreateSalesOrder   string
	CreateSalesInvoice string

	UpdateCustomerCard string
	UpdateItemCard     string
	UpdateSalesOrder   string
	UpdateSalesInvoice string
}

var Query = QueryType{
	AssemblyBom:  "AssemblyBom",
	CustomerCard: "CustomerCard",
	ItemCard:     "ItemCard",
	SalesOrder:   "SalesOrder",
	SalesInvoice: "SalesInvoice",
}

var Mutation = MutationType{
	CreateCustomerCard: "CreateCustomerCard",
	CreateItemCard:     "CreateItemCard",
	CreateSalesOrder:   "CreateSalesOrder",
	CreateSalesInvoice: "CreateSalesInvoice",

	UpdateCustomerCard: "UpdateCustomerCard",
	UpdateItemCard:     "UpdateItemCard",
	UpdateSalesOrder:   "UpdateSalesOrder",
	UpdateSalesInvoice: "UpdateSalesInvoice",
}
