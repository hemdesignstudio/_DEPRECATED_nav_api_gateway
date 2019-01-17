package errorhandler

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

func errorOccurredButNotUnmarshalled() error {
	return fmt.Errorf("error occured, however Gateway API has failed to unmarshal the error response from NAV")
}

func BadRequest(statusCode int, handler Handler) error {
	return fmt.Errorf("status: %d, %s, %s", statusCode, handler.Content.Code, handler.Content.Message)
}

func NavServerError(statusCode int, handler Handler) error {
	return fmt.Errorf("status: %d, %s, %s", statusCode, handler.Content.Code, handler.Content.Message)
}

func ValueIsNotCorrect(args map[string]interface{}) error {
	return fmt.Errorf(" there are no entries for '%s' of value '%s' ", args["key"], args["value"])
}

func CouldNotUnmarshalData() error {
	return fmt.Errorf("could not unmarshal data")
}

func unauthorized() error {
	return fmt.Errorf("NAV credentials in GateWay configuration is not correct")

}

func Handle(statusCode int, errBody []byte) error {
	handler := Handler{}
	err := json.Unmarshal(errBody, &handler)
	if err != nil {
		return errorOccurredButNotUnmarshalled()
	}

	if statusCode == http.StatusBadRequest {
		return BadRequest(statusCode, handler)
	}

	if statusCode == http.StatusUnauthorized {
		return unauthorized()
	}

	return NavServerError(statusCode, handler)
}
