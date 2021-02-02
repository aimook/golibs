package v

type H map[string]interface{}

func Success() H {
	return H{
		"success": true,
		"message": "success",
		"code":    10000,
	}
}

func Fail(code int) H {
	return H{
		"success": false,
		"message": "fail",
		"code":    code,
	}
}

func (h H) Message(message string) H {
	h["message"] = message
	return h
}

func (h H) Data(data interface{}) H {
	h["data"] = data
	return h
}
