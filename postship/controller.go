// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package postship

import (
	"github.com/hem-nav-gateway/request"
)

//GetAll retrieves a List of all Posted Sales Shipments available Microsoft Navision .
//Function takes a list of fields to be returned by Microsoft Navision.
func GetAll(fields interface{}) (interface{}, error) {

	return request.GetAll(endpoint, fields, Response{})

}

//Filter retrieves a list of filtered Posted Sales Shipments based on a key-value pair added by the requester
//Function takes a list of fields to be returned by Microsoft Navision.
//Function takes filter arguments which are required for filtering results in Navision
func Filter(fields, args interface{}) (interface{}, error) {

	return request.Filter(endpoint, fields, args, Response{})

}
