// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package salesinvoice

import (
	"github.com/hem-nav-gateway/request"
)

//GetAll retrieves a List of all SalesInvoices available Microsoft Navision .
//Function takes a list of fields to be returned by Microsoft Navision.
//SalesInvoice has nested GraphQL fields SalesLine.
//SalesLine is removed is it is only known to this API and should not be sent to Navision
func GetAll(fields interface{}) (interface{}, error) {
	fields = removeField("Sales_Lines", fields)
	return request.GetAll(endpoint, fields, Response{})
}

//Filter retrieves a list of filtered SalesInvoices based on a key-value pair added by the requester
//Function takes a list of fields to be returned by Microsoft Navision.
//Function takes filter arguments which are required for filtering results in Navision
//SalesInvoice has nested GraphQL fields SalesLine.
//SalesLine is removed is it is only known to this API and should not be sent to Navision
func Filter(fields, args interface{}) (interface{}, error) {
	fields = removeField("Sales_Lines", fields)
	return request.Filter(endpoint, fields, args, Response{})
}

//Create creates a SalesInvoice objects based on arguments added by the requester
//Function takes a list of fields to be returned by Microsoft Navision after creation.
//SalesInvoice has nested GraphQL fields SalesLine.
//SalesLine is removed is it is only known to this API and should not be sent to Navision
func Create(fields, args interface{}) (interface{}, error) {
	fields = removeField("Sales_Lines", fields)
	return request.Create(endpoint, fields, args, Response{})
}

//Update modifies a certain SalesInvoice Object Microsoft Navision.
//Function takes filter arguments which are required identifying
//the specific object to be updated/modified.
//Function returns a list of AssemblyBom values
//For SalesInvoice, Navision requires an extra argument
//Which Is document Type in this case it is Document_Type = "Invoice"
//SalesInvoice has nested GraphQL fields SalesLine.
//SalesLine is removed is it is only known to this API and should not be sent to Navision
func Update(fields, args interface{}) (interface{}, error) {
	docType := "Invoice"
	fields = removeField("Sales_Lines", fields)
	return request.Update(endpoint, fields, args, docType, Response{})
}
