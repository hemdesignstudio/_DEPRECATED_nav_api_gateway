package salesline

import "github.com/hem-nav-gateway/request"

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
	_args := args.(map[string]interface{})
	docType := _args["Document_Type"]
	return request.Update(endpoint, args, docType, Response{})
}
