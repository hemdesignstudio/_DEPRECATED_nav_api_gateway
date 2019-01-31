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

func Update(endpoint string, args map[string]interface{}, response interface{}) (interface{}, error) {
	no := args["No"].(string)
	body, _ := json.Marshal(args)
	resByte, resError := updateEntitybyId(endpoint, no, body)
	if resError != nil {
		return nil, resError
	}
	err := json.Unmarshal(resByte, &response)
	if err != nil {
		return nil, errorhandler.CouldNotUnmarshalData()
	}
	return response, nil
}
