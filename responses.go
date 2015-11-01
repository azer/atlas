package atlas

var (
	NotFound      = Error(404, "404 - Not found.")
	InternalError = Error(500, "500 - Internal Server Error")
)

func Manual(code int, anything interface{}) *Response {
	return &Response{code, anything, nil, true}
}

func Success(anything interface{}) *Response {
	return &Response{
		200,
		&SuccessResponseContext{anything, true},
		nil,
		true,
	}
}

func Error(code int, err interface{}) *Response {
	return &Response{
		code,
		&ErrorResponseContext{err},
		nil,
		true,
	}
}

func Custom(output []byte) *Response {
	return &Response{
		200,
		nil,
		output,
		false,
	}
}
