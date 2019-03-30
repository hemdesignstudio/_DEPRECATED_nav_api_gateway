// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package rest

type utils struct{}

func newRequest() *restClient {
	return &restClient{}
}

// get handles all requests with http method "GET"
func (*utils) get(uri string) (int, interface{}, error) {
	r := newRequest()
	return r.request(uri, "GET", nil)
}

// post handles all requests with http method "POST"
func (*utils) post(uri string, body interface{}) (int, interface{}, error) {
	r := newRequest()
	return r.request(uri, "POST", body)

}

// patch handles all requests with http method "PATCH"
func (*utils) patch(uri string, body interface{}) (int, interface{}, error) {
	r := newRequest()
	return r.request(uri, "PATCH", body)
}

// delete handles all requests with http method "DELETE"
func (*utils) delete(uri string, body interface{}) (int, interface{}, error) {
	r := newRequest()
	return r.request(uri, "DELETE", body)
}
