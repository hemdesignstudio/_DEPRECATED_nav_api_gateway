package errorhandler

type Handler struct {
	Message string  `json:"message"`
	Content Content `json:"error"`
}

type Content struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}
