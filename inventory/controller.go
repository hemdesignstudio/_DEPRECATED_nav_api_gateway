package inventory

import (
	"github.com/graphql-go/graphql"
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

// GetAll retrieves a List of all Inventory Items available Microsoft Navision .
// Function takes a list of fields to be returned by Microsoft Navision.
func (r *Request) GetAll() (interface{}, error) {
	return soapRequest(getPayload(nil))
}

func (r *Request) Filter() (interface{}, error) {
	return soapRequest(getPayload(r.Object.Args))
}

func (r *Request) Create() (interface{}, error) {
	return nil, nil
}

func (r *Request) Update() (interface{}, error) {
	return nil, nil
}
