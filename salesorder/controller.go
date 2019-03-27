// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package salesorder

import (
	"github.com/graphql-go/graphql"
	"github.com/hem-nav-gateway/request"
	"github.com/hem-nav-gateway/types"
)

var _type = createType()
var _args = createArgs()

type Request struct {
	Company string
	Object  types.RequestObject
}

func (*Request) CreateType() *graphql.Object {
	return _type
}

func (*Request) CreateArgs() map[string]*graphql.ArgumentConfig {
	return _args
}

func (r *Request) GetCompany() string {
	return r.Company
}

func (r *Request) SetArgs(args map[string]interface{}) {
	r.Object.Args = args
}

func (r *Request) SetFields(fields []string) {
	r.Object.Fields = fields
}

// GetAll retrieves a List of all SalesOrders available Microsoft Navision .
// Function takes a list of fields to be returned by Microsoft Navision.
func (r *Request) GetAll() (interface{}, error) {

	// SalesOrder has nested GraphQL field SalesLine.
	// SalesLine is removed is it is only known to this API
	// and should not be sent to Navision
	r.Object.Endpoint = endpoint
	r.Object.Company = r.Company
	_salesLine.Company = r.Company
	r.Object.Fields = removeField("Sales_Lines", r.Object.Fields)

	r.Object.Response = Response{}

	return request.GetAll(r.Object)
}

// Filter retrieves a list of filtered SalesOrders
// based on a key-value pair added by the requester.
// Function takes a list of fields to be returned by Microsoft Navision.
// Function takes filter arguments which are
// required for filtering results in Navision.
func (r *Request) Filter() (interface{}, error) {

	// SalesOrder has nested GraphQL field SalesLine.
	// SalesLine is removed is it is only known to this API
	// and should not be sent to Navision
	r.Object.Endpoint = endpoint
	r.Object.Company = r.Company
	_salesLine.Company = r.Company

	r.Object.Fields = removeField("Sales_Lines", r.Object.Fields)

	r.Object.Response = Response{}

	return request.Filter(r.Object)
}

// Create creates a SalesOrder objects based on arguments added by the requester
// Function takes a list of fields to be returned by Microsoft Navision after creation.
// Function takes filter arguments which are
// required for creating a new object
func (r *Request) Create() (interface{}, error) {

	// SalesOrder has nested GraphQL field SalesLine.
	// SalesLine is removed is it is only known to this API
	// and should not be sent to Navision
	r.Object.Endpoint = endpoint
	r.Object.Company = r.Company
	_salesLine.Company = r.Company

	r.Object.Fields = removeField("Sales_Lines", r.Object.Fields)

	r.Object.Response = Response{}

	return request.Create(r.Object)
}

// Update modifies a certain SalesOrder Object Microsoft Navision.
// Function takes arguments which are required identifying
// the specific object to be updated/modified.
func (r *Request) Update() (interface{}, error) {

	// For SalesOrder, Navision requires an extra argument
	// which Is document Type in this case it is Document_Type = "Order".

	// SalesOrder has nested GraphQL field SalesLine.
	// SalesLine is removed is it is only known to this API
	// and should not be sent to Navision
	r.Object.Endpoint = endpoint
	r.Object.Company = r.Company
	_salesLine.Company = r.Company

	r.Object.Properties = map[string]interface{}{}
	r.Object.Properties["docType"] = "Order"

	r.Object.Fields = removeField("Sales_Lines", r.Object.Fields)

	r.Object.Response = Response{}

	return request.Update(r.Object)
}
