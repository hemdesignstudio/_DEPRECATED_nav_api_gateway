// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package assemblybom

import (
	"github.com/hem-nav-gateway/request"
)

//GetAll takes a list of fields to be returned by Microsoft Navision.
//Function returns a list of AssemblyBom values
func GetAll(fields interface{}) (interface{}, error) {

	return request.GetAll(endpoint, fields, response{})

}

//Filter takes a list of fields to be returned by Microsoft Navision.
//Function takes filter arguments which are required for filtering results in Navision
//Function returns a list of AssemblyBom values
func Filter(fields, args interface{}) (interface{}, error) {

	return request.Filter(endpoint, fields, args, response{})
}
