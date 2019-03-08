// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package item

import (
	"github.com/hem-nav-gateway/request"
	"github.com/hem-nav-gateway/types"
)

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
func Create(object interface{}) (interface{}, error) {
	obj := object.(types.RequestObject)
	obj.Endpoint = endpoint
	return request.Create(obj, Response{})
}

// Update modifies a certain ItemCard Object Microsoft Navision.
// Function takes filter arguments which are required identifying
// the specific object to be updated/modified.
// Function returns a list of AssemblyBom values.
func Update(object interface{}) (interface{}, error) {
	obj := object.(types.RequestObject)
	obj.Endpoint = endpoint
	return request.Update(obj, Response{})
}
