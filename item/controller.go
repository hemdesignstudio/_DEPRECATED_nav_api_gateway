package item

import (
	"github.com/hem-nav-gateway/request"
)

func GetAll(fields interface{}) (interface{}, error) {
	return request.GetAll(endpoint, fields, Response{})
}

func Filter(fields, args interface{}) (interface{}, error) {
	return request.Filter(endpoint, fields, args, Response{})
}

func Create(fields, args interface{}) (interface{}, error) {
	return request.Create(endpoint, fields, args, Response{})
}

func Update(fields, args interface{}) (interface{}, error) {
	return request.Update(endpoint, fields, args, nil, Response{})
}
