// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package salesinvoice

import (
	"github.com/hem-nav-gateway/request"
)

// GetAll retrieves a List of all SalesInvoices available Microsoft Navision .
// Function takes a list of fields to be returned by Microsoft Navision.
func GetAll(fields interface{}) (interface{}, error) {

	// SalesInvoice has nested GraphQL field SalesLine.
	// SalesLine is removed is it is only known to this API
	// and should not be sent to Navision
	fields = removeField("Sales_Lines", fields)

	return request.GetAll(endpoint, fields, Response{})
}

// Filter retrieves a list of filtered SalesInvoices
// based on a key-value pair added by the requester.
// Function takes a list of fields to be returned by Microsoft Navision.
// Function takes filter arguments which are
// required for filtering results in Navision.
func Filter(fields, args interface{}) (interface{}, error) {

	// SalesInvoice has nested GraphQL field SalesLine.
	// SalesLine is removed is it is only known to this API
	// and should not be sent to Navision
	fields = removeField("Sales_Lines", fields)

	return request.Filter(endpoint, fields, args, Response{})
}

// Create creates a SalesInvoice objects based on arguments added by the requester
// Function takes a list of fields to be returned by Microsoft Navision after creation.
// Function takes filter arguments which are
// required for creating a new object
func Create(fields, args interface{}) (interface{}, error) {

	// SalesInvoice has nested GraphQL field SalesLine.
	// SalesLine is removed is it is only known to this API and should not be sent to Navision
	fields = removeField("Sales_Lines", fields)

	return request.Create(endpoint, fields, args, Response{})
}

// Update modifies a certain SalesInvoice Object Microsoft Navision.
// Function takes arguments which are required identifying
// the specific object to be updated/modified.
// Function returns a list of AssemblyBom values
func Update(fields, args interface{}) (interface{}, error) {

	// For SalesInvoice, Navision requires an extra argument
	// which Is document Type in this case it is Document_Type = "Invoice".
	docType := "Invoice"

	// SalesInvoice has nested GraphQL field SalesLine.
	// SalesLine is removed is it is only known to this API
	// and should not be sent to Navision
	fields = removeField("Sales_Lines", fields)

	return request.Update(endpoint, fields, args, docType, Response{})
}
