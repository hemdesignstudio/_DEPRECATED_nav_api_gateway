// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package errorhandler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// badRequest is invoked when Microsft Navision returns response code "400"
// returns response code, Navision code and message to requester
func badRequest(statusCode int, handler Handler) error {
	return fmt.Errorf("status: %d, %s, %s", statusCode, handler.Content.Code, handler.Content.Message)
}

// errorMessageFromNav is invoked when the response code from Microsft Navision != 200
// returns response code, Navision code and message to requester
func errorMessageFromNav(statusCode int, handler Handler) error {
	return fmt.Errorf("status: %d, %s, %s", statusCode, handler.Content.Code, handler.Content.Message)
}

// navInternalServerError is invoked when Microsft Navision returns response code "500"
// returns response code, Navision code and message to requester
func navInternalServerError(statusCode int, handler Handler) error {
	return fmt.Errorf("status: %d, %s, Internal Server Error in Microsoft NAV", statusCode, handler.Message)
}

// ValueIsNotCorrect is invoked when requester, uses incorrect filter arguments
// returns wrong arguments to the requester
func ValueIsNotCorrect(args interface{}) error {
	_args := args.(map[string]interface{})
	return fmt.Errorf(" there are no entries for '%s' of value '%s' ", _args["key"], _args["value"])
}

// CouldNotUnmarshalData in invoked when The API Gateway
// in unable to decode JSON from Navision
func CouldNotUnmarshalData() error {
	return fmt.Errorf("could not unmarshal data, NAV API Gateway internal server error")
}

/*
errorOccurredButNotUnmarshalled is invoked when Microsoft Navision returns
error which does not follow the normal format.

Normal error response format:

	'''
	{
		"error": {
			"code": "BadRequest",
			"message": "Could not find a property named 'hello' on type 'NAV.CustomerCardWS'."
		}
	}
	'''

However sometimes navision returns this:

	'''
	{
		"Message": "An error has occurred."
	}
	'''

 Or other formats and atm. this Gateway API handles only the first type
*/
func errorOccurredButNotUnmarshalled() error {
	return fmt.Errorf("error occured, however Gateway API has failed to unmarshal the error response from NAV")
}

// unauthorized is invoked when Microsft Navision returns response code "401"
func unauthorized() error {
	return fmt.Errorf("NAV credentials in GateWay configuration is not correct")

}

// CompanyDoesNotExist is invoked when wrong company is requested in URL path
func CompanyDoesNotExist(name string) error {
	return fmt.Errorf(fmt.Sprintf(" '/%s' does no exist, available options are '/test' ", name))

}

// Handle is the main function for handling all types of errors
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
