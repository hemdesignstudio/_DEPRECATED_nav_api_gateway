// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package request

import (
	"fmt"
	"github.com/hem-nav-gateway/config"
	"github.com/hem-nav-gateway/types"
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

// resolveBaseUrl resolves the base URL for endpoint and company name
func resolveBaseUrl(endpoint, companyName string) string {

	uri := config.BaseUrl +
		config.CompanyEndpoint +
		fmt.Sprintf("('%v')", companyName) +
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
func getAllEntities(object interface{}) interface{} {

	obj := object.(types.RequestObject)
	baseUri := resolveBaseUrl(obj.Endpoint, obj.Company)
	returnFields := resolveFields(obj.Fields)
	uri := baseUri + "?" + returnFields

	_, respBody, _ := get(uri)
	return respBody
}

// createEntity creates a specific entity
func createEntity(object, body interface{}) (interface{}, error) {

	obj := object.(types.RequestObject)
	baseUri := resolveBaseUrl(obj.Endpoint, obj.Company)
	returnFields := resolveFields(obj.Fields)
	uri := baseUri + "?" + returnFields

	_, respBody, err := post(uri, body)
	return respBody, err

}

// filterByArgs gets a list of entities based on a filter
func filterByArgs(object interface{}) (interface{}, error) {

	obj := object.(types.RequestObject)
	baseUri := resolveBaseUrl(obj.Endpoint, obj.Company)
	filter := resolveFilterArgs(obj.Args)
	returnFields := resolveFields(obj.Fields)
	uri := baseUri + "?" + filter + "&" + returnFields

	_, respBody, err := get(uri)
	return respBody, err
}

// updateEntitybyId updates a specific entity
// takes the unique id of the entity
// and fields to be returned after update
func updateEntitybyId(object, body interface{}) (interface{}, error) {

	obj := object.(types.RequestObject)
	baseUri := resolveBaseUrl(obj.Endpoint, obj.Company)
	returnFields := resolveFields(obj.Fields)

	selector := fmt.Sprintf(
		"('%v')",
		obj.Args["No"],
	)

	uri := baseUri + selector + "?" + returnFields

	_, respBody, err := patch(uri, body)
	return respBody, err

}

// updateEntitybyDocumentTypeAndID updates a specific entity
// takes the unique id and document_type of the entity
// and fields to be returned after update
func updateEntitybyDocumentTypeAndID(object, body interface{}) (interface{}, error) {

	obj := object.(types.RequestObject)
	baseUri := resolveBaseUrl(obj.Endpoint, obj.Company)
	returnFields := resolveFields(obj.Fields)

	selector := fmt.Sprintf(
		"('%v','%v')",
		obj.Properties["docType"],
		obj.Args["No"],
	)

	uri := baseUri + selector + "?" + returnFields

	_, respBody, err := patch(uri, body)
	return respBody, err

}

// updateEntitybyDocumentTypeAndIDAndLineNo updates a specific entity, mostly utilized for updating SalesLines
// takes the unique id, document_type and Line_no of the entity
// and fields to be returned after update
func updateEntitybyDocumentTypeAndIDAndLineNo(object interface{}, body interface{}) (interface{}, error) {

	obj := object.(types.RequestObject)
	baseUri := resolveBaseUrl(obj.Endpoint, obj.Company)
	returnFields := resolveFields(obj.Fields)

	selector := fmt.Sprintf(
		"('%v','%v',%v)",
		obj.Properties["docType"],
		obj.Args["Document_No"],
		obj.Args["Line_No"],
	)

	uri := baseUri + selector + "?" + returnFields

	_, respBody, err := patch(uri, body)
	return respBody, err

}

// deleteEntitybyId deletes a specific entity
// takes the unique id of the entity
// and fields to be returned after update
func deleteEntitybyId(endpoint, id string) (int, error) {
	baseUri := resolveBaseUrl(endpoint, config.TestCompany)
	selector := fmt.Sprintf("('%s')", id)
	uri := baseUri + selector
	resCode, _, err := delete(uri, nil)
	return resCode, err
}

// deleteEntitybyDocumentTypeAndID delets a specific entity
// takes the unique id and document_type of the entity
// and fields to be returned after update
func deleteEntitybyDocumentTypeAndID(endpoint, id, docType string) (int, error) {
	baseUri := resolveBaseUrl(endpoint, config.TestCompany)
	selector := fmt.Sprintf("('%s','%s')", docType, id)
	uri := baseUri + selector
	resCode, _, err := delete(uri, nil)
	return resCode, err

}

// deleteEntitybyDocumentTypeAndIDAndLineNo deletes a specific entity, mostly utilized for updating SalesLines
// takes the unique id, document_type and Line_no of the entity
// and fields to be returned after update
func deleteEntitybyDocumentTypeAndIDAndLineNo(endpoint, id, docType string, lineNo int) (int, error) {
	baseUri := resolveBaseUrl(endpoint, config.TestCompany)
	selector := fmt.Sprintf("('%s','%s',%d)", docType, id, lineNo)
	uri := baseUri + selector
	resCode, _, err := delete(uri, nil)
	return resCode, err

}
