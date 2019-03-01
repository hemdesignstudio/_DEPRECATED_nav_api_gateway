package request

import (
	"encoding/json"
	"github.com/hem-nav-gateway/errorhandler"
)

func GetAll(endpoint string, fields, response interface{}) (interface{}, error) {
	resByte := getAllEntities(endpoint, fields)
	err := json.Unmarshal(resByte, &response)
	if err != nil {
		return nil, errorhandler.CouldNotUnmarshalData()
	}
	res := response.(map[string]interface{})
	return res["value"], nil
}

func Filter(endpoint string, fields, args, response interface{}) (interface{}, error) {
	resByte, resError := filterByArgs(endpoint, fields, args)
	if resError != nil {
		return nil, resError
	}
	err := json.Unmarshal(resByte, &response)
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

func Create(endpoint string, fields, args, response interface{}) (interface{}, error) {
	body, _ := json.Marshal(args)
	resByte, resError := createEntity(endpoint, fields, body)
	if resError != nil {
		return nil, resError
	}
	err := json.Unmarshal(resByte, &response)
	if err != nil {
		return nil, errorhandler.CouldNotUnmarshalData()
	}
	return response, nil
}

func Update(endpoint string, fields, args, docType, response interface{}) (interface{}, error) {

	var resByte []byte
	var resError error

	body, _ := json.Marshal(args)
	_args := args.(map[string]interface{})
	if docType != nil {
		docType := docType.(string)
		if lineNo, ok := _args["Line_No"]; ok {
			id := _args["Document_No"].(string)
			lineNo := lineNo.(int)
			resByte, resError = updateEntitybyDocumentTypeAndIDAndLineNo(endpoint, id, docType, fields, lineNo, body)
		} else {
			id := _args["No"].(string)
			resByte, resError = updateEntitybyDocumentTypeAndID(endpoint, id, docType, fields, body)
		}
	} else {
		id := _args["No"].(string)
		resByte, resError = updateEntitybyId(endpoint, id, fields, body)
	}

	if resError != nil {
		return nil, resError
	}
	err := json.Unmarshal(resByte, &response)
	if err != nil {
		return nil, errorhandler.CouldNotUnmarshalData()
	}
	return response, nil
}

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
