// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package assemblybom

import (
	"github.com/graphql-go/graphql"
	"github.com/hem-nav-gateway/request"
	"github.com/hem-nav-gateway/types"
)

var _type = createType()

type Request struct {
	Company string
	Object  types.RequestObject
}

func (*Request) CreateType() *graphql.Object {
	return _type
}

func (*Request) CreateArgs() map[string]*graphql.ArgumentConfig {
	return nil
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

// GetAll retrieves all of AssemblyBom objects in Microsoft Navision.
// Function takes a list of AssemblyBom fields to be returned
func (r *Request) GetAll() (interface{}, error) {

	r.Object.Endpoint = endpoint
	r.Object.Company = r.Company
	return request.GetAll(r.Object, response{})

}

// Filter retrieves a filtered list of AssemblyBom objects based of key-value pair.
// Funtion takes a list of fields to be returned by Microsoft Navision.
// Function takes filter arguments which are required for filtering results in Navision
func (r *Request) Filter() (interface{}, error) {
	r.Object.Endpoint = endpoint
	r.Object.Company = r.Company
	return request.Filter(r.Object, response{})
}

func (r *Request) Create() (interface{}, error) {
	return nil, nil
}

func (r *Request) Update() (interface{}, error) {
	return nil, nil
}
