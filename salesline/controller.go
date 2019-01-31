package salesline

import "github.com/nav-api-gateway/request"

func GetAll() (interface{}, error) {
	res := Response{}
	return request.GetAll(endpoint, res)
}

func Filter(args map[string]interface{}) (interface{}, error) {
	res := Response{}
	return request.Filter(endpoint, args, res)
}

func Create(args map[string]interface{}) (interface{}, error) {
	res := Response{}
	return request.Create(endpoint, args, res)
}

func Update(args map[string]interface{}) (interface{}, error) {
	res := Response{}
	return request.Update(endpoint, args, nil, res)
}
