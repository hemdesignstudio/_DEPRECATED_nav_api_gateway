package postship

import (
	"github.com/nav-api-gateway/request"
)

func GetAll() (interface{}, error) {
	res := Response{}
	return request.GetAll(companyName, endpoint, res)
}

func Filter(args map[string]interface{}) (interface{}, error) {
	res := Response{}
	return request.Filter(companyName, endpoint, args, res)
}

func Create(args map[string]interface{}) (interface{}, error) {
	res := Response{}
	return request.Create(companyName, endpoint, args, res)
}

func Update(args map[string]interface{}) (interface{}, error) {
	res := Response{}
	return request.Update(companyName, endpoint, args, res)
}
