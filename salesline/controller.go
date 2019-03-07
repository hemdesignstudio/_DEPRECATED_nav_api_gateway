// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package salesline

import (
	"github.com/hem-nav-gateway/request"
	"github.com/hem-nav-gateway/types"
)

type RequestObject struct {
	Endpoint string
	Company  string
	Fields   interface{}
	Args     map[string]interface{}
}

// GetAll retrieves a List of all ItemCards available Microsoft Navision .
// Function takes a list of fields to be returned by Microsoft Navision.
func GetAll(object interface{}) (interface{}, error) {
	obj := object.(types.RequestObject)
	obj.Endpoint = endpoint
	return request.GetAll(obj, Response{})
}

// Filter retrieves a list of filtered ItemCards based on a key-value pair added by the requester
// Function takes a list of fields to be returned by Microsoft Navision.
// Function takes filter arguments which are required for filtering results in Navision.
func Filter(object interface{}) (interface{}, error) {
	obj := object.(types.RequestObject)
	obj.Endpoint = endpoint
	return request.Filter(obj, Response{})
}

// Create creates a ItemCard objects based on arguments  added by the requester.
// Function takes a list of fields to be returned by Microsoft Navision after creation.
func Create(fields, args interface{}) (interface{}, error) {
	return request.Create(endpoint, fields, args, Response{})
}

// Update modifies a certain ItemCard Object Microsoft Navision.
// Function takes filter arguments which are required identifying
// the specific object to be updated/modified.
// Function returns a list of AssemblyBom values.
func Update(fields, args interface{}) (interface{}, error) {

	// SalesLines in Microsoft Navision are divided into two types
	// - Document_Type = 'Order' - for SalesLines which are based on SalesOrder.
	// - Document_Type = 'Invoice' - for SalesLines which are related on SalesInvoice.
	// It is a requirement to specify type for Navision Odata REST API incase of
	// updating an entity
	_args := args.(map[string]interface{})
	docType := _args["Document_Type"]
	return request.Update(endpoint, fields, args, docType, Response{})
}
