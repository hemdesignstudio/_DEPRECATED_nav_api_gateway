// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package request

import (
	"encoding/json"
	"github.com/hem-nav-gateway/errorhandler"
	"github.com/hem-nav-gateway/types"
)

// GetAll handles getting all entities for a specified object types (customer, item ... etc)
// takes fields --> fields to be returned from Navision
// returns a list of requests object values
func GetAll(object, response interface{}) (interface{}, error) {
	obj := object.(types.RequestObject)
	resValue := getAllEntities(obj.Endpoint, obj.Fields)
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
func Filter(object, response interface{}) (interface{}, error) {
	obj := object.(types.RequestObject)

	resValue, resError := filterByArgs(obj.Endpoint, obj.Fields, obj.Args)
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
		return nil, errorhandler.ValueIsNotCorrect(obj.Args)
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

		// In Order to update SalesLines for example
		// It is required to specify "Line_No", "Document_type", "Document_No"
		// In this specific case "Document_No" acts as id
		// this is related to how Microsoft Navision works
		if lineNo, ok := _args["Line_No"]; ok {
			id := _args["Document_No"].(string)
			lineNo := lineNo.(int)
			resValue, resError = updateEntitybyDocumentTypeAndIDAndLineNo(endpoint, id, docType, fields, lineNo, body)

			// In order to update SalesOrder or SalesInvoice
			// it is required to specify its "No" which acts as its id
			// and "Document_type" which specifies if it is "Order" or "Invoice"
		} else {
			id := _args["No"].(string)
			resValue, resError = updateEntitybyDocumentTypeAndID(endpoint, id, docType, fields, body)
		}

		// This is the case for most entities to be updated
		// "No" which acts as id is all what it required to update an entity
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

		// In Order to delete and entity like "SalesLines" for example
		// It is required to specify "Line_No", "Document_type", "Document_No"
		// In this specific case "Document_No" acts as id
		// this is related to how Microsoft Navision works
		if lineNo, ok := args["Line_No"]; ok {
			id := args["Document_No"].(string)
			lineNo := lineNo.(int)
			resCode, resError = deleteEntitybyDocumentTypeAndIDAndLineNo(endpoint, id, docType, lineNo)

			// In order to delete SalesOrder or SalesInvoice
			// it is required to specify its "No" which acts as its id
			// and "Document_type" which specifies if it is "Order" or "Invoice"
		} else {
			id := args["No"].(string)
			resCode, resError = deleteEntitybyDocumentTypeAndID(endpoint, id, docType)
		}

		// This is the case for most entities to be deleted
		// "No" which acts as id is all what it required to delete an entity
	} else {
		id := args["No"].(string)
		resCode, resError = deleteEntitybyId(endpoint, id)
	}

	if resError != nil {
		return nil, resError
	}

	return resCode, nil
}
