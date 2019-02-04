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

	_, respBody, _ := get(uri)
	return respBody
}

func createEntity(endpoint string, body []byte) ([]byte, error) {
	uri := config.BaseUrl +
		config.CompanyEndpoint +
		fmt.Sprintf("('%s')", config.CompanyName) +
		endpoint

	_, respBody, err := post(uri, body)
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

	_, respBody, err := get(uri + filter)
	return respBody, err
}

func updateEntitybyId(endpoint, id string, body []byte) ([]byte, error) {
	uri := config.BaseUrl +
		config.CompanyEndpoint +
		fmt.Sprintf("('%s')", config.CompanyName) +
		endpoint + fmt.Sprintf("('%s')", id)

	_, respBody, err := patch(uri, body)
	return respBody, err

}

func updateEntitybyDocumentTypeAndID(endpoint, id, docType string, body []byte) ([]byte, error) {
	uri := config.BaseUrl +
		config.CompanyEndpoint +
		fmt.Sprintf("('%s')", config.CompanyName) +
		endpoint + fmt.Sprintf("('%s','%s')", docType, id)

	_, respBody, err := patch(uri, body)
	return respBody, err

}

func updateEntitybyDocumentTypeAndIDAndLineNo(endpoint, id, docType string, lineNo int, body []byte) ([]byte, error) {
	uri := config.BaseUrl +
		config.CompanyEndpoint +
		fmt.Sprintf("('%s')", config.CompanyName) +
		endpoint + fmt.Sprintf("('%s','%s',%d)", docType, id, lineNo)

	_, respBody, err := patch(uri, body)
	return respBody, err

}

func deleteEntitybyId(endpoint, id string) (int, error) {
	uri := config.BaseUrl +
		config.CompanyEndpoint +
		fmt.Sprintf("('%s')", config.CompanyName) +
		endpoint + fmt.Sprintf("('%s')", id)

	resCode, _, err := delete(uri, nil)
	return resCode, err
}

func deleteEntitybyDocumentTypeAndID(endpoint, id, docType string) (int, error) {
	uri := config.BaseUrl +
		config.CompanyEndpoint +
		fmt.Sprintf("('%s')", config.CompanyName) +
		endpoint + fmt.Sprintf("('%s','%s')", docType, id)

	resCode, _, err := delete(uri, nil)
	return resCode, err

}

func deleteEntitybyDocumentTypeAndIDAndLineNo(endpoint, id, docType string, lineNo int) (int, error) {
	uri := config.BaseUrl +
		config.CompanyEndpoint +
		fmt.Sprintf("('%s')", config.CompanyName) +
		endpoint + fmt.Sprintf("('%s','%s',%d)", docType, id, lineNo)

	resCode, _, err := delete(uri, nil)
	return resCode, err

}
