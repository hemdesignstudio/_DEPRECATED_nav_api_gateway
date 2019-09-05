// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package customer

import (
	"github.com/graphql-go/graphql"
	"github.com/hem-nav-gateway/rest"
	"github.com/hem-nav-gateway/types"
)

var _type = createType()
var _args = createArgs()

type Request struct {
	Company string
	Object  types.RequestObject
}

func newRestService() *rest.Service {
	return &rest.Service{}
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

// GetAll retrieves a List of all CustomerCards available Microsoft Navision .
// Function takes a list of fields to be returned by Microsoft Navision.
func (r *Request) GetAll() (interface{}, error) {
	r.Object.Endpoint = endpoint
	r.Object.Company = r.Company
	r.Object.Response = Response{}
	s := newRestService()
	return s.GetAll(r.Object)
}

// Filter retrieves a list of filtered CustomerCards based on a key-value pair added by the requester
// Function takes a list of fields to be returned by Microsoft Navision.
// Function takes filter arguments which are required for filtering results in Navision
func (r *Request) Filter() (interface{}, error) {
	r.Object.Endpoint = endpoint
	r.Object.Company = r.Company
	r.Object.Response = Response{}
	s := newRestService()

	return s.Filter(r.Object)

}

// Create creates a CustomerCard objects based on arguments added by the requester
// Function takes a list of fields to be returned by Microsoft Navision after creation.
func (r *Request) Create() (interface{}, error) {
	r.Object.Endpoint = endpoint
	r.Object.Company = r.Company
	r.Object.Response = Response{}
	s := newRestService()

	return s.Create(r.Object)

}

// Update modifies a certain CustomerCard Object Microsoft Navision.
// Function takes filter arguments which are required identifying
// the specific object to be updated/modified.
// Function returns a list of AssemblyBom values
func (r *Request) Update() (interface{}, error) {
	r.Object.Endpoint = endpoint
	r.Object.Company = r.Company
	r.Object.Response = Response{}
	s := newRestService()

	return s.Update(r.Object)

}
