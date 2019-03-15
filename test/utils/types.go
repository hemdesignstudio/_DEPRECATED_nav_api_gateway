package utils

type ErrorMessage struct {
	Message interface{} `json:"message"`
}

type QueryType struct {
	AssemblyBom  string
	CustomerCard string
	Inventory    string
	ItemCard     string
	SalesOrder   string
	SalesInvoice string
	PostShip     string
	SalesLine    string
}

type MutationType struct {
	CreateCustomerCard string
	CreateItemCard     string
	CreateSalesOrder   string
	CreateSalesInvoice string
	CreateSalesLine    string

	UpdateCustomerCard string
	UpdateItemCard     string
	UpdateSalesOrder   string
	UpdateSalesInvoice string
	UpdateSalesLine    string
}

var Query = QueryType{
	AssemblyBom:  "AssemblyBom",
	CustomerCard: "CustomerCard",
	Inventory:    "Inventory",
	ItemCard:     "ItemCard",
	SalesOrder:   "SalesOrder",
	SalesInvoice: "SalesInvoice",
	PostShip:     "PostedSalesShipment",
	SalesLine:    "SalesLine",
}

var Mutation = MutationType{
	CreateCustomerCard: "CreateCustomerCard",
	CreateItemCard:     "CreateItemCard",
	CreateSalesOrder:   "CreateSalesOrder",
	CreateSalesInvoice: "CreateSalesInvoice",
	CreateSalesLine:    "CreateSalesLine",

	UpdateCustomerCard: "UpdateCustomerCard",
	UpdateItemCard:     "UpdateItemCard",
	UpdateSalesOrder:   "UpdateSalesOrder",
	UpdateSalesInvoice: "UpdateSalesInvoice",
	UpdateSalesLine:    "UpdateSalesLine",
}
