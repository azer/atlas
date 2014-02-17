package atlas

import (
	"encoding/json"
	"errors"
)

type Response struct {
	Code    int
	Context interface{}
}

type SuccessResponseContext struct {
	Result interface{} `json:"result"`
	Ok     bool        `json:"ok"`
}

type ErrorResponseContext struct {
	Error interface{} `json:"error"`
}

var NotFound = Error(404, errors.New("404 - Not found."))

func (response *Response) Stringify() string {
	parsed, _ := json.MarshalIndent(response.Context, "", "	")
	return string(parsed)
}

func Manual(code int, anything interface{}) *Response {
	return &Response{code, anything}
}

func Success(anything interface{}) *Response {
	return &Response{
		200,
		&SuccessResponseContext{anything, true},
	}
}

func Error(code int, err interface{}) *Response {
	return &Response{
		code,
		&ErrorResponseContext{err},
	}
}
