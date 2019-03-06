// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package request

import (
	"fmt"
	"github.com/hem-nav-gateway/config"
	"reflect"
	"strings"
)

// resolveFields creates the URL query for
// return fields from Microsoft Navision
// Function takes fields --> []string
// Function returns the select query for Microsoft Navision
// Example :
// 			fields = ["No", "Name", "Address"]
// 			Return "$select=No,Name,Adress"
func resolveFields(fields interface{}) string {
	if fields == nil {
		return ""
	}
	fieldList := fields.([]string)
	returnFields := fmt.Sprintf("$select=%v", strings.Join(fieldList, ", "))
	return returnFields
}

// resolveBaseUrl resolves the base URL for endpoint
func resolveBaseUrl(endpoint string) string {

	uri := config.BaseUrl +
		config.CompanyEndpoint +
		fmt.Sprintf("('%v')", config.TestCompanyName) +
		endpoint
	return uri
}

// resolveFilterArgs creates URL query used filtering
// Function takes args --> map[string]interface{}
// Function returns the filter query for Microsoft Navision
// Example :
// 			args = {"key":"No", "value":"1001"}
// 			Return "$filter=No eq '1001'"
// Hint: Microsoft Navision filtering understands types as follows
// $filter=No eq '1001' --> type of '1001' is string since single quotes are used
// $filter=No eq 1001 --> type of 1001 is integer or number since no quotes are used
// Navision does not accept double quotes
// $filter=No eq "1001" --> does not work
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

// getAllEntities gets all required entities
func getAllEntities(endpoint string, fields interface{}) interface{} {

	baseUri := resolveBaseUrl(endpoint)
	returnFields := resolveFields(fields)
	uri := baseUri + "?" + returnFields

	_, respBody, _ := get(uri)
	return respBody
}

// createEntity creates a specific entity
func createEntity(endpoint string, fields, body interface{}) (interface{}, error) {

	baseUri := resolveBaseUrl(endpoint)
	returnFields := resolveFields(fields)
	uri := baseUri + "?" + returnFields

	_, respBody, err := post(uri, body)
	return respBody, err

}

// filterByArgs gets a list of entities based on a filter
func filterByArgs(endpoint string, fields, args interface{}) (interface{}, error) {

	baseUri := resolveBaseUrl(endpoint)
	filter := resolveFilterArgs(args)
	returnFields := resolveFields(fields)
	uri := baseUri + "?" + filter + "&" + returnFields

	_, respBody, err := get(uri)
	return respBody, err
}

// updateEntitybyId updates a specific entity
// takes the unique id of the entity
// and fields to be returned after update
func updateEntitybyId(endpoint, id string, fields, body interface{}) (interface{}, error) {

	baseUri := resolveBaseUrl(endpoint)
	selector := fmt.Sprintf("('%s')", id)
	returnFields := resolveFields(fields)
	uri := baseUri + selector + "?" + returnFields

	_, respBody, err := patch(uri, body)
	return respBody, err

}

// updateEntitybyDocumentTypeAndID updates a specific entity
// takes the unique id and document_type of the entity
// and fields to be returned after update
func updateEntitybyDocumentTypeAndID(endpoint, id, docType string, fields, body interface{}) (interface{}, error) {

	baseUri := resolveBaseUrl(endpoint)
	selector := fmt.Sprintf("('%s','%s')", docType, id)
	returnFields := resolveFields(fields)
	uri := baseUri + selector + "?" + returnFields

	_, respBody, err := patch(uri, body)
	return respBody, err

}

// updateEntitybyDocumentTypeAndIDAndLineNo updates a specific entity, mostly utilized for updating SalesLines
// takes the unique id, document_type and Line_no of the entity
// and fields to be returned after update
func updateEntitybyDocumentTypeAndIDAndLineNo(endpoint, id, docType string, fields interface{}, lineNo int, body interface{}) (interface{}, error) {
	baseUri := resolveBaseUrl(endpoint)
	selector := fmt.Sprintf("('%s','%s',%d)", docType, id, lineNo)
	returnFields := resolveFields(fields)
	uri := baseUri + selector + "?" + returnFields

	_, respBody, err := patch(uri, body)
	return respBody, err

}

// deleteEntitybyId deletes a specific entity
// takes the unique id of the entity
// and fields to be returned after update
func deleteEntitybyId(endpoint, id string) (int, error) {
	uri := config.BaseUrl +
		config.CompanyEndpoint +
		fmt.Sprintf("('%v')", config.TestCompanyName) +
		endpoint + fmt.Sprintf("('%s')", id)

	resCode, _, err := delete(uri, nil)
	return resCode, err
}

// deleteEntitybyDocumentTypeAndID delets a specific entity
// takes the unique id and document_type of the entity
// and fields to be returned after update
func deleteEntitybyDocumentTypeAndID(endpoint, id, docType string) (int, error) {
	uri := config.BaseUrl +
		config.CompanyEndpoint +
		fmt.Sprintf("('%v')", config.TestCompanyName) +
		endpoint + fmt.Sprintf("('%s','%s')", docType, id)

	resCode, _, err := delete(uri, nil)
	return resCode, err

}

// deleteEntitybyDocumentTypeAndIDAndLineNo deletes a specific entity, mostly utilized for updating SalesLines
// takes the unique id, document_type and Line_no of the entity
// and fields to be returned after update
func deleteEntitybyDocumentTypeAndIDAndLineNo(endpoint, id, docType string, lineNo int) (int, error) {
	uri := config.BaseUrl +
		config.CompanyEndpoint +
		fmt.Sprintf("('%v')", config.TestCompanyName) +
		endpoint + fmt.Sprintf("('%s','%s',%d)", docType, id, lineNo)

	resCode, _, err := delete(uri, nil)
	return resCode, err

}
