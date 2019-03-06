// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package request

import (
	"fmt"
	"github.com/hem-nav-gateway/config"
	"github.com/hem-nav-gateway/session"
	"reflect"
	"strings"
)

func resolveFields(fields interface{}) string {
	if fields == nil {
		return ""
	}
	fieldList := fields.([]string)
	returnFields := fmt.Sprintf("$select=%v", strings.Join(fieldList, ", "))
	return returnFields
}

func resolveBaseUrl(endpoint string) string {

	uri := config.BaseUrl +
		config.CompanyEndpoint +
		fmt.Sprintf("('%v')", session.CompanyName()) +
		endpoint
	return uri
}

func resolveFilterArgs(args interface{}) string {

	_args := args.(map[string]interface{})

	key := _args["key"]
	value := _args["value"]
	valueType := reflect.TypeOf(value).Kind()

	filter := fmt.Sprintf("$filter=%v eq %v", key, value)

	if valueType == reflect.String {
		filter = fmt.Sprintf("$filter=%v eq '%v'", key, value)
	}
	return filter

}

func getAllEntities(endpoint string, fields interface{}) interface{} {

	baseUri := resolveBaseUrl(endpoint)
	returnFields := resolveFields(fields)
	uri := baseUri + "?" + returnFields

	_, respBody, _ := get(uri)
	return respBody
}

func createEntity(endpoint string, fields, body interface{}) (interface{}, error) {

	baseUri := resolveBaseUrl(endpoint)
	returnFields := resolveFields(fields)
	uri := baseUri + "?" + returnFields

	_, respBody, err := post(uri, body)
	return respBody, err

}

func filterByArgs(endpoint string, fields, args interface{}) (interface{}, error) {

	baseUri := resolveBaseUrl(endpoint)
	filter := resolveFilterArgs(args)
	returnFields := resolveFields(fields)
	uri := baseUri + "?" + filter + "&" + returnFields

	_, respBody, err := get(uri)
	return respBody, err
}

func updateEntitybyId(endpoint, id string, fields, body interface{}) (interface{}, error) {

	baseUri := resolveBaseUrl(endpoint)
	selector := fmt.Sprintf("('%s')", id)
	returnFields := resolveFields(fields)
	uri := baseUri + selector + "?" + returnFields

	_, respBody, err := patch(uri, body)
	return respBody, err

}

func updateEntitybyDocumentTypeAndID(endpoint, id, docType string, fields, body interface{}) (interface{}, error) {

	baseUri := resolveBaseUrl(endpoint)
	selector := fmt.Sprintf("('%s','%s')", docType, id)
	returnFields := resolveFields(fields)
	uri := baseUri + selector + "?" + returnFields

	_, respBody, err := patch(uri, body)
	return respBody, err

}

func updateEntitybyDocumentTypeAndIDAndLineNo(endpoint, id, docType string, fields interface{}, lineNo int, body interface{}) (interface{}, error) {
	baseUri := resolveBaseUrl(endpoint)
	selector := fmt.Sprintf("('%s','%s',%d)", docType, id, lineNo)
	returnFields := resolveFields(fields)
	uri := baseUri + selector + "?" + returnFields

	_, respBody, err := patch(uri, body)
	return respBody, err

}

func deleteEntitybyId(endpoint, id string) (int, error) {
	uri := config.BaseUrl +
		config.CompanyEndpoint +
		fmt.Sprintf("('%v')", session.CompanyName()) +
		endpoint + fmt.Sprintf("('%s')", id)

	resCode, _, err := delete(uri, nil)
	return resCode, err
}

func deleteEntitybyDocumentTypeAndID(endpoint, id, docType string) (int, error) {
	uri := config.BaseUrl +
		config.CompanyEndpoint +
		fmt.Sprintf("('%v')", session.CompanyName()) +
		endpoint + fmt.Sprintf("('%s','%s')", docType, id)

	resCode, _, err := delete(uri, nil)
	return resCode, err

}

func deleteEntitybyDocumentTypeAndIDAndLineNo(endpoint, id, docType string, lineNo int) (int, error) {
	uri := config.BaseUrl +
		config.CompanyEndpoint +
		fmt.Sprintf("('%v')", session.CompanyName()) +
		endpoint + fmt.Sprintf("('%s','%s',%d)", docType, id, lineNo)

	resCode, _, err := delete(uri, nil)
	return resCode, err

}
