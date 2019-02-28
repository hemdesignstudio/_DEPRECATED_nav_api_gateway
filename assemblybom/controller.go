package assemblybom

import (
	"github.com/hem-nav-gateway/config"
	"github.com/hem-nav-gateway/request"
)

var endpoint = config.AssemblyBomEndpoint

func GetAll() (interface{}, error) {
	res := Response{}
	return request.GetAll(endpoint, res)
}

func Filter(args map[string]interface{}) (interface{}, error) {
	res := Response{}
	return request.Filter(endpoint, args, res)
}
