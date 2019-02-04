package request

import (
	"encoding/json"
	"github.com/nav-api-gateway/errorhandler"
)

func GetAll(endpoint string, response interface{}) (interface{}, error) {
	resByte := getAllEntities(endpoint)
	err := json.Unmarshal(resByte, &response)
	if err != nil {
		return nil, errorhandler.CouldNotUnmarshalData()
	}
	res := response.(map[string]interface{})
	return res["value"], nil
}

func Filter(endpoint string, args map[string]interface{}, response interface{}) (interface{}, error) {
	resByte, resError := filterByArgs(endpoint, args)
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

func Create(endpoint string, args map[string]interface{}, response interface{}) (interface{}, error) {
	body, _ := json.Marshal(args)
	resByte, resError := createEntity(endpoint, body)
	if resError != nil {
		return nil, resError
	}
	err := json.Unmarshal(resByte, &response)
	if err != nil {
		return nil, errorhandler.CouldNotUnmarshalData()
	}
	return response, nil
}

func Update(endpoint string, args map[string]interface{}, docType, response interface{}) (interface{}, error) {

	var resByte []byte
	var resError error

	body, _ := json.Marshal(args)

	if docType != nil {
		docType := docType.(string)
		if lineNo, ok := args["Line_No"]; ok {
			id := args["Document_No"].(string)
			lineNo := lineNo.(int)
			resByte, resError = updateEntitybyDocumentTypeAndIDAndLineNo(endpoint, id, docType, lineNo, body)
		} else {
			id := args["No"].(string)
			resByte, resError = updateEntitybyDocumentTypeAndID(endpoint, id, docType, body)
		}
	} else {
		id := args["No"].(string)
		resByte, resError = updateEntitybyId(endpoint, id, body)
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
