package errorhandler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func badRequest(statusCode int, handler Handler) error {
	return fmt.Errorf("status: %d, %s, %s", statusCode, handler.Content.Code, handler.Content.Message)
}

func errorMessageFromNav(statusCode int, handler Handler) error {
	return fmt.Errorf("status: %d, %s, %s", statusCode, handler.Content.Code, handler.Content.Message)
}

func navInternalServerError(statusCode int, handler Handler) error {
	return fmt.Errorf("status: %d, %s, Internal Server Error in Microsoft NAV", statusCode, handler.Message)
}

func ValueIsNotCorrect(args map[string]interface{}) error {
	return fmt.Errorf(" there are no entries for '%s' of value '%s' ", args["key"], args["value"])
}

func CouldNotUnmarshalData() error {
	return fmt.Errorf("could not unmarshal data")
}

func errorOccurredButNotUnmarshalled() error {
	return fmt.Errorf("error occured, however Gateway API has failed to unmarshal the error response from NAV")
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
		return badRequest(statusCode, handler)
	}

	if statusCode == http.StatusUnauthorized {
		return unauthorized()
	}

	if statusCode == http.StatusInternalServerError {
		return navInternalServerError(statusCode, handler)
	}

	return errorMessageFromNav(statusCode, handler)
}
