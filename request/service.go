// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package request

import (
	"encoding/json"
	"github.com/hem-nav-gateway/errorhandler"
)

// GetAll handles getting all entities for a specified object types (customer, item ... etc)
// takes fields --> fields to be returned from Navision
// returns a list of requests object values
func GetAll(endpoint string, fields, response interface{}) (interface{}, error) {
	resValue := getAllEntities(endpoint, fields)
	err := json.Unmarshal(resValue.([]byte), &response)
	if err != nil {
		return nil, errorhandler.CouldNotUnmarshalData()
	}
	res := response.(map[string]interface{})
	return res["value"], nil
}

// Filter handles getting specific entities of specified object types (customer, item ... etc)
// takes fields --> fields to be returned from Navision
// takes args --> filter arguments
// returns a list of filtered object values
func Filter(endpoint string, fields, args, response interface{}) (interface{}, error) {
	resValue, resError := filterByArgs(endpoint, fields, args)
	if resError != nil {
		return nil, resError
	}
	err := json.Unmarshal(resValue.([]byte), &response)
	if err != nil {
		return nil, errorhandler.CouldNotUnmarshalData()
	}
	res := response.(map[string]interface{})
	values := res["value"].([]interface{})
	if len(values) == 0 {
		return nil, errorhandler.ValueIsNotCorrect(args)
	}
	return values, nil
}

// Create a specific entity of a specified object type (customer, item ... etc)
// takes fields --> fields to be returned from Navision
// takes args --> arguments are object values to be created
// returns the created object with its return fields
func Create(endpoint string, fields, args, response interface{}) (interface{}, error) {
	body, _ := json.Marshal(args)
	resValue, resError := createEntity(endpoint, fields, body)
	if resError != nil {
		return nil, resError
	}
	err := json.Unmarshal(resValue.([]byte), &response)
	if err != nil {
		return nil, errorhandler.CouldNotUnmarshalData()
	}
	return response, nil
}

// Update a specific entity of a specified object type (customer, item ... etc)
// takes fields --> fields to be returned from Navision
// takes args --> arguments are object values to be updated
// returns the created object with its return fields
func Update(endpoint string, fields, args, docType, response interface{}) (interface{}, error) {

	var resValue interface{}
	var resError error

	body, _ := json.Marshal(args)
	_args := args.(map[string]interface{})
	if docType != nil {
		docType := docType.(string)
		if lineNo, ok := _args["Line_No"]; ok {
			id := _args["Document_No"].(string)
			lineNo := lineNo.(int)
			resValue, resError = updateEntitybyDocumentTypeAndIDAndLineNo(endpoint, id, docType, fields, lineNo, body)
		} else {
			id := _args["No"].(string)
			resValue, resError = updateEntitybyDocumentTypeAndID(endpoint, id, docType, fields, body)
		}
	} else {
		id := _args["No"].(string)
		resValue, resError = updateEntitybyId(endpoint, id, fields, body)
	}

	if resError != nil {
		return nil, resError
	}
	err := json.Unmarshal(resValue.([]byte), &response)
	if err != nil {
		return nil, errorhandler.CouldNotUnmarshalData()
	}
	return response, nil
}

// Delete a specific entity of a specified object type (customer, item ... etc)
// takes args --> arguments used to get entity to be deleted
// returns the response code
func Delete(endpoint string, args map[string]interface{}, docType interface{}) (interface{}, error) {

	var resCode int
	var resError error

	if docType != nil {
		docType := docType.(string)
		if lineNo, ok := args["Line_No"]; ok {
			id := args["Document_No"].(string)
			lineNo := lineNo.(int)
			resCode, resError = deleteEntitybyDocumentTypeAndIDAndLineNo(endpoint, id, docType, lineNo)
		} else {
			id := args["No"].(string)
			resCode, resError = deleteEntitybyDocumentTypeAndID(endpoint, id, docType)
		}
	} else {
		id := args["No"].(string)
		resCode, resError = deleteEntitybyId(endpoint, id)
	}

	if resError != nil {
		return nil, resError
	}

	return resCode, nil
}
