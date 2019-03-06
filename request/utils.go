// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package request

// get handles all requests with http method "GET"
func get(uri string) (int, interface{}, error) {
	return request(uri, "GET", nil)
}

// post handles all requests with http method "POST"
func post(uri string, body interface{}) (int, interface{}, error) {
	return request(uri, "POST", body)

}

// patch handles all requests with http method "PATCH"
func patch(uri string, body interface{}) (int, interface{}, error) {
	return request(uri, "PATCH", body)
}

// delete handles all requests with http method "DELETE"
func delete(uri string, body interface{}) (int, interface{}, error) {
	return request(uri, "DELETE", body)
}
