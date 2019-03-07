package fields

import (
	"github.com/hem-nav-gateway/assemblybom"
	"github.com/hem-nav-gateway/customer"
	"github.com/hem-nav-gateway/item"
	"github.com/hem-nav-gateway/postship"
	"github.com/hem-nav-gateway/salesinvoice"
	"github.com/hem-nav-gateway/salesline"
	"github.com/hem-nav-gateway/salesorder"
)

func assemblyBomField(company string) fieldType {
	field := fieldType{
		Name:    "assemblyBom",
		Company: company,
		Filter:  assemblybom.Filter,
		GetAll:  assemblybom.GetAll,
	}
	return field
}

func customerField(company string) fieldType {
	field := fieldType{
		Name:    "customer",
		Company: company,
		Filter:  customer.Filter,
		GetAll:  customer.GetAll,
		Create:  customer.Create,
		Update:  customer.Update,
	}
	return field
}

func itemField(company string) fieldType {
	field := fieldType{
		Name:    "item",
		Company: company,
		Filter:  item.Filter,
		GetAll:  item.GetAll,
		Create:  item.Create,
		Update:  item.Update,
	}
	return field
}

func salesOrderField(company string) fieldType {
	field := fieldType{
		Name:    "salesOrder",
		Company: company,
		Filter:  salesorder.Filter,
		GetAll:  salesorder.GetAll,
		Create:  salesorder.Create,
		Update:  salesorder.Update,
	}
	return field
}

func salesLineField(company string) fieldType {
	field := fieldType{
		Name:    "salesLine",
		Company: company,
		Filter:  salesline.Filter,
		GetAll:  salesline.GetAll,
		Create:  salesline.Create,
		Update:  salesline.Update,
	}
	return field
}

func postShipField(company string) fieldType {
	field := fieldType{
		Name:    "postShip",
		Company: company,
		Filter:  postship.Filter,
		GetAll:  postship.GetAll,
	}
	return field
}

func salesInvoiceField(company string) fieldType {
	field := fieldType{
		Name:    "salesInvoice",
		Company: company,
		Filter:  salesinvoice.Filter,
		GetAll:  salesinvoice.GetAll,
		Create:  salesinvoice.Create,
		Update:  salesinvoice.Update,
	}
	return field
}
