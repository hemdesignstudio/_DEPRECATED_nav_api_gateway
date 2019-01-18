package request

func get(uri string) ([]byte, error) {
	return request(uri, "GET", nil)
}

func post(uri string, body []byte) ([]byte, error) {
	return request(uri, "POST", body)

}

func patch(uri string, body []byte) ([]byte, error) {
	return request(uri, "PATCH", body)
}
