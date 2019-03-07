// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package assemblybom

import (
	"github.com/hem-nav-gateway/request"
	"github.com/hem-nav-gateway/types"
)

// GetAll retrieves all of AssemblyBom objects in Microsoft Navision.
// Function takes a list of AssemblyBom fields to be returned
func GetAll(object interface{}) (interface{}, error) {

	obj := object.(types.RequestObject)
	obj.Endpoint = endpoint
	return request.GetAll(obj, response{})

}

// Filter retrieves a filtered list of AssemblyBom objects based of key-value pair.
// Funtion takes a list of fields to be returned by Microsoft Navision.
// Function takes filter arguments which are required for filtering results in Navision
func Filter(object interface{}) (interface{}, error) {
	obj := object.(types.RequestObject)
	obj.Endpoint = endpoint
	return request.Filter(obj, response{})
}
