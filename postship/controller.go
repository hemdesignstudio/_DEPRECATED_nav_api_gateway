package postship

import (
	"github.com/hem-nav-gateway/request"
)

func GetAll() (interface{}, error) {
	return request.GetAll(endpoint, Response{})
}

func Filter(args interface{}) (interface{}, error) {
	return request.Filter(endpoint, args, Response{})
}
