package request

import (
	"encoding/json"
	"github.com/nav-api-gateway/errorhandler"
)

func GetAllByCompanyName(companyName string, endpoint string, response interface{}) (interface{}, error) {
	resByte := GetAll(companyName, endpoint)
	err := json.Unmarshal(resByte, &response)
	if err != nil {
		return nil, errorhandler.CouldNotUnmarshalData()
	}
	res := response.(map[string]interface{})
	return res["value"], nil
}

func FilterByArgs(companyName, endpoint string, args map[string]interface{}, response interface{}) (interface{}, error) {
	resByte, resError := Filter(companyName, endpoint, args)
	if resError != nil {
		return nil, resError
	}
	err := json.Unmarshal(resByte, &response)
	if err != nil {
		return nil, errorhandler.CouldNotUnmarshalData()
	}
	res := response.(map[string]interface{})

	if len(res) == 0 {
		return nil, errorhandler.ValueIsNotCorrect(args)
	}
	return res["value"], nil
}
