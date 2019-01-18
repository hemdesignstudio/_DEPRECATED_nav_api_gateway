package request

import (
	"fmt"
	"github.com/nav-api-gateway/config"
	"reflect"
)

func getAllEntitiesByCompanyName(companyName string, endpoint string) []byte {
	uri := config.BaseUrl +
		config.CompanyEndpoint +
		fmt.Sprintf("('%s')", companyName) +
		endpoint

	respBody, _ := get(uri)
	return respBody
}

func filterByArgs(companyName, endpoint string, args map[string]interface{}) ([]byte, error) {
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

	respBody, err := get(uri + filter)
	return respBody, err
}

func createEntity(companyName string, endpoint string, body []byte) ([]byte, error) {
	uri := config.BaseUrl +
		config.CompanyEndpoint +
		fmt.Sprintf("('%s')", companyName) +
		endpoint

	respBody, err := post(uri, body)
	return respBody, err

}

func updateEntitybyId(companyName string, endpoint string, id string, body []byte) ([]byte, error) {
	uri := config.BaseUrl +
		config.CompanyEndpoint +
		fmt.Sprintf("('%s')", companyName) +
		endpoint + fmt.Sprintf("('%s')", id)

	respBody, err := patch(uri, body)
	return respBody, err

}
