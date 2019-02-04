package request

func get(uri string) (int, []byte, error) {
	return request(uri, "GET", nil)
}

func post(uri string, body []byte) (int, []byte, error) {
	return request(uri, "POST", body)

}

func patch(uri string, body []byte) (int, []byte, error) {
	return request(uri, "PATCH", body)
}

func delete(uri string, body []byte) (int, []byte, error) {
	return request(uri, "DELETE", body)
}
