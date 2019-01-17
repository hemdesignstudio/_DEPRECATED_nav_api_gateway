package assemblybom

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
