package error

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler struct {
	Content Content `json:"error"`
}

type Content struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func ErrorOccuredButNotUnmarshaled() error {
	return fmt.Errorf("error occured, however Gateway API has failed to unmarshal the error Response from NAV")
}

func BadRequest(statusCode int, handler Handler) error {
	return fmt.Errorf("status: %d, %s, %s", statusCode, handler.Content.Code, handler.Content.Message)
}

func UnknownError() error {
	return fmt.Errorf("there has been an unknown urror which has occured")
}

func ValueIsNotCorrect(args map[string]interface{}) error {
	return fmt.Errorf(" there are no entries for '%s' of value '%s' ", args["key"], args["value"])
}

func Handle(statusCode int, errBody []byte) error {
	handler := Handler{}
	err := json.Unmarshal(errBody, &handler)
	if err != nil {
		return ErrorOccuredButNotUnmarshaled()
	}

	if statusCode == http.StatusBadRequest {
		return BadRequest(statusCode, handler)
	}

	return UnknownError()
}
