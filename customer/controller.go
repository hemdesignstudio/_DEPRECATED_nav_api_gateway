package customer

import (
	"github.com/hem-nav-gateway/request"
)

func GetAll() (interface{}, error) {

	return request.GetAll(endpoint, Response{})
}

func Filter(args map[string]interface{}) (interface{}, error) {

	return request.Filter(endpoint, args, Response{})
}

func Create(args map[string]interface{}) (interface{}, error) {

	return request.Create(endpoint, args, Response{})
}

func Update(args map[string]interface{}) (interface{}, error) {

	return request.Update(endpoint, args, nil, Response{})
}
