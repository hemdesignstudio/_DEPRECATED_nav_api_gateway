package request

import (
	"fmt"
	"github.com/nav-api-gateway/config"
	"reflect"
)

func getAllEntities(endpoint string) []byte {
	uri := config.BaseUrl +
		config.CompanyEndpoint +
		fmt.Sprintf("('%s')", config.CompanyName) +
		endpoint

	respBody, _ := get(uri)
	return respBody
}

func createEntity(endpoint string, body []byte) ([]byte, error) {
	uri := config.BaseUrl +
		config.CompanyEndpoint +
		fmt.Sprintf("('%s')", config.CompanyName) +
		endpoint

	respBody, err := post(uri, body)
	return respBody, err

}

func filterByArgs(endpoint string, args map[string]interface{}) ([]byte, error) {
	uri := config.BaseUrl +
		config.CompanyEndpoint +
		fmt.Sprintf("('%s')", config.CompanyName) +
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

func updateEntitybyId(endpoint, id string, body []byte) ([]byte, error) {
	uri := config.BaseUrl +
		config.CompanyEndpoint +
		fmt.Sprintf("('%s')", config.CompanyName) +
		endpoint + fmt.Sprintf("('%s')", id)

	respBody, err := patch(uri, body)
	return respBody, err

}

func updateEntitybyDocumentTypeAndID(endpoint, id, docType string, body []byte) ([]byte, error) {
	uri := config.BaseUrl +
		config.CompanyEndpoint +
		fmt.Sprintf("('%s')", config.CompanyName) +
		endpoint + fmt.Sprintf("('%s','%s')", docType, id)

	respBody, err := patch(uri, body)
	return respBody, err

}

func updateEntitybyDocumentTypeAndIDAndLineNo(endpoint, id, docType string, lineNo int, body []byte) ([]byte, error) {
	uri := config.BaseUrl +
		config.CompanyEndpoint +
		fmt.Sprintf("('%s')", config.CompanyName) +
		endpoint + fmt.Sprintf("('%s','%s',%d)", docType, id, lineNo)

	respBody, err := patch(uri, body)
	return respBody, err

}
