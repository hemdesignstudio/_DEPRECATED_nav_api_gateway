// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package salesinvoice

import (
	"github.com/hem-nav-gateway/request"
	"github.com/hem-nav-gateway/types"
)

// GetAll retrieves a List of all SalesInvoices available Microsoft Navision .
// Function takes a list of fields to be returned by Microsoft Navision.
func GetAll(object interface{}) (interface{}, error) {

	// SalesInvoice has nested GraphQL field SalesLine.
	// SalesLine is removed is it is only known to this API
	// and should not be sent to Navision
	obj := object.(types.RequestObject)
	obj.Endpoint = endpoint
	obj.Fields = removeField("Sales_Lines", obj.Fields)

	return request.GetAll(obj, Response{})
}

// Filter retrieves a list of filtered SalesInvoices
// based on a key-value pair added by the requester.
// Function takes a list of fields to be returned by Microsoft Navision.
// Function takes filter arguments which are
// required for filtering results in Navision.
func Filter(object interface{}) (interface{}, error) {

	// SalesInvoice has nested GraphQL field SalesLine.
	// SalesLine is removed is it is only known to this API
	// and should not be sent to Navision
	obj := object.(types.RequestObject)
	obj.Endpoint = endpoint
	obj.Fields = removeField("Sales_Lines", obj.Fields)

	return request.Filter(obj, Response{})
}

// Create creates a SalesInvoice objects based on arguments added by the requester
// Function takes a list of fields to be returned by Microsoft Navision after creation.
// Function takes filter arguments which are
// required for creating a new object
func Create(object interface{}) (interface{}, error) {

	// SalesInvoice has nested GraphQL field SalesLine.
	// SalesLine is removed is it is only known to this API and should not be sent to Navision
	obj := object.(types.RequestObject)
	obj.Endpoint = endpoint
	obj.Fields = removeField("Sales_Lines", obj.Fields)

	return request.Create(obj, Response{})
}

// Update modifies a certain SalesInvoice Object Microsoft Navision.
// Function takes arguments which are required identifying
// the specific object to be updated/modified.
// Function returns a list of AssemblyBom values
func Update(object interface{}) (interface{}, error) {

	// For SalesInvoice, Navision requires an extra argument
	// which Is document Type in this case it is Document_Type = "Invoice".

	// SalesInvoice has nested GraphQL field SalesLine.
	// SalesLine is removed is it is only known to this API
	// and should not be sent to Navision
	obj := object.(types.RequestObject)
	obj.Properties = map[string]interface{}{}
	obj.Properties["docType"] = "Invoice"
	obj.Endpoint = endpoint
	obj.Fields = removeField("Sales_Lines", obj.Fields)

	return request.Update(obj, Response{})
}
