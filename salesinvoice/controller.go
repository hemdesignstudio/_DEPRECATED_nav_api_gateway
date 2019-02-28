package salesinvoice

import (
	"github.com/hem-nav-gateway/request"
)

func GetAll() (interface{}, error) {

	return request.GetAll(endpoint, Response{})
}

func Filter(args interface{}) (interface{}, error) {

	return request.Filter(endpoint, args, Response{})
}

func Create(args interface{}) (interface{}, error) {

	return request.Create(endpoint, args, Response{})
}

func Update(args interface{}) (interface{}, error) {
	docType := "Invoice"
	return request.Update(endpoint, args, docType, Response{})
}
