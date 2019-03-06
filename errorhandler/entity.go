// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

/*
Package errorhandler implements a simple package for handling errors related
to Hem Nav API Gateway

Handler type defined in this package is used to decode JSON coming from Microsft Navision In case of
errors

Example Error Response from Microsoft Navision:

	'''
	{
		"error": {
			"code": "BadRequest",
			"message": "Could not find a property named 'hello' on type 'NAV.CustomerCardWS'."
		}
	}
	'''

Example way of decoding JSON error message
	handler := Handler{}
	err := json.Unmarshal(errorBody, &handler)

*/
package errorhandler

type Handler struct {
	Message string  `json:"message"`
	Content Content `json:"error"`
}

type Content struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}
