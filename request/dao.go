package request

import (
	"fmt"
	"github.com/nav-api-gateway/config"
	"reflect"
)

func GetAll(companyName string, endpoint string) []byte {
	uri := config.BaseUrl +
		config.CompanyEndpoint +
		fmt.Sprintf("('%s')", companyName) +
		endpoint

	resultByte, _ := GET(uri)
	return resultByte
}

func Filter(companyName, endpoint string, args map[string]interface{}) ([]byte, error) {
	uri := config.BaseUrl +
		config.CompanyEndpoint +
		fmt.Sprintf("('%s')", companyName) +
		endpoint

	key := args["key"]
	value := args["value"]
	valueType := reflect.TypeOf(args["value"]).Kind()
	filter := fmt.Sprintf("?$filter=%s eq %s", key, value)

	if valueType == reflect.String {
		filter = fmt.Sprintf("?$filter=%s eq '%s'", key, value)
	}

	resultByte, err := GET(uri + filter)
	return resultByte, err
}

func Create(companyName string, endpoint string, body []byte) []byte {
	uri := config.BaseUrl +
		config.CompanyEndpoint +
		fmt.Sprintf("('%s')", companyName) +
		endpoint

	resultByte, _ := POST(uri, body)
	return resultByte

}

func Update(companyName string, endpoint string, id string, body []byte) []byte {
	uri := config.BaseUrl +
		config.CompanyEndpoint +
		fmt.Sprintf("('%s')", companyName) +
		endpoint + fmt.Sprintf("('%s')", id)

	resultByte, _ := PATCH(uri, body)

	return resultByte

}
