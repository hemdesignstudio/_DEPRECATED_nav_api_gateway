package assemblybom

import (
	"github.com/hem-nav-gateway/request"
)

func GetAll(fields interface{}) (interface{}, error) {

	return request.GetAll(endpoint, fields, response{})

}

func Filter(fields, args interface{}) (interface{}, error) {

	return request.Filter(endpoint, fields, args, response{})
}
