// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package salesinvoice

import (
	"github.com/hem-nav-gateway/request"
)

func GetAll(fields interface{}) (interface{}, error) {
	fields = removeField("Sales_Lines", fields)
	return request.GetAll(endpoint, fields, Response{})
}

func Filter(fields, args interface{}) (interface{}, error) {
	fields = removeField("Sales_Lines", fields)
	return request.Filter(endpoint, fields, args, Response{})
}

func Create(fields, args interface{}) (interface{}, error) {
	fields = removeField("Sales_Lines", fields)
	return request.Create(endpoint, fields, args, Response{})
}

func Update(fields, args interface{}) (interface{}, error) {
	docType := "Invoice"
	fields = removeField("Sales_Lines", fields)
	return request.Update(endpoint, fields, args, docType, Response{})
}
