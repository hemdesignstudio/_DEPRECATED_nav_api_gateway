package assemblybom

import (
	"github.com/hem-nav-gateway/request"
)

type Request struct {
	endpoint string
	fields   interface{}
	args     interface{}
	response *response
}

func GetAll(fields interface{}) (interface{}, error) {

	return request.GetAll(endpoint, fields, response{})

}

func Filter(fields, args interface{}) (interface{}, error) {

	return request.Filter(endpoint, fields, args, response{})
}
